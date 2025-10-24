package test

import (
	"context"
	"os"
	"testing"
	"time"

	"net/http/httptest"

	"github.com/gin-gonic/gin"

	db "github.com/chocobone/articode_web/db/config"
	"github.com/chocobone/articode_web/user"
	userRepo "github.com/chocobone/articode_web/user/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// helper: try to connect to local MongoDB, skip test if unavailable
func connectOrSkip(t *testing.T) *mongo.Client {
	t.Helper()
	uri := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Skipf("skipping test; cannot connect to mongo: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(context.Background())
		t.Skipf("skipping test; cannot ping mongo: %v", err)
	}

	// assign into package-level vars used by the app
	db.Client = client
	db.UserCollection = client.Database("resq").Collection("users")
	db.ModelingCollection = client.Database("resq").Collection("3D_Modeling")

	return client
}

func TestDBConnectionAndUserRepositoryCRUD(t *testing.T) {
	client := connectOrSkip(t)
	defer func() {
		_ = client.Disconnect(context.Background())
	}()

	// Basic sanity checks for exported collections
	if db.UserCollection == nil {
		t.Fatal("UserCollection is nil after connecting to mongo")
	}
	if db.ModelingCollection == nil {
		t.Fatal("ModelingCollection is nil after connecting to mongo")
	}

	// repository integration: create -> read -> delete
	repo := userRepo.NewUserRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testID := int(time.Now().Unix() % 1000000) // simple unique-ish id
	u := &userRepo.UserInfoResponse{
		UserID: testID,
		Name:   "integration-test",
		Email:  "it@example.test",
	}

	// insert
	created, err := repo.PostUserInfo(ctx, u)
	if err != nil {
		t.Fatalf("PostUserInfo failed: %v", err)
	}
	if created == nil || created.UserID != testID {
		t.Fatalf("unexpected created user: %+v", created)
	}

	// fetch
	got, err := repo.GetUserInfo(ctx, testID)
	if err != nil {
		t.Fatalf("GetUserInfo failed: %v", err)
	}
	if got.Name != u.Name || got.Email != u.Email {
		t.Fatalf("retrieved user mismatch: got %+v want %+v", got, u)
	}

	// delete
	if err := repo.DeleteUserInfo(ctx, testID); err != nil {
		t.Fatalf("DeleteUserInfo failed: %v", err)
	}

	// ensure deleted
	_, err = repo.GetUserInfo(ctx, testID)
	if err == nil {
		t.Fatalf("expected GetUserInfo to fail after delete")
	}
}

func TestRouterHealthRoute(t *testing.T) {
	// this test only verifies router setup; it doesn't bind network ports
	gin.SetMode(gin.TestMode)

	// create repo/service/handler similarly to main
	client := connectOrSkip(t)
	defer func() { _ = client.Disconnect(context.Background()) }()

	userRepoInst := userRepo.NewUserRepository()
	userSvc := user.NewUserService(userRepoInst)
	userHandler := user.NewUserHandler(userSvc)

	r := gin.New()
	// register minimal middleware used by app (CORS + options handling)
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// health route (replicate main.StatusHandler)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running!"})
	})

	// mount user routes (this verifies route registration doesn't panic)
	// note: GetUserRoutes is expected to exist and accept these args
	user.GetUserRoutes(r, userHandler)

	// exercise health endpoint
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("health route returned status %d, want 200", w.Code)
	}
}
