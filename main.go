package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"                   //gin framework
	"github.com/joho/godotenv"                   //local variable on .env files
	httpSwagger "github.com/swaggo/http-swagger" //auto-API docs writter

	db "github.com/chocobone/articode_web/db/config"

	// ---- User API ---- //
	"github.com/chocobone/articode_web/user"
	userRepository "github.com/chocobone/articode_web/user/repository"

	// ---- 3D API ---- //
	"github.com/chocobone/articode_web/modeling3d"
	modelingRepository "github.com/chocobone/articode_web/modeling3d/repository"
)

func StatusHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running!",
	})
}

var (
	host = "localhost"
	port = "6100"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR: .env file NOT FOUND")
	}

	os.Getenv("CLIENT_ID")
	fmt.Println("CLIENT_SECRET: test") //, os.Getenv("CLIENT_SECRET"))

	db.InitMongo()

	r := gin.Default()
	r.Static("/static", "./")

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// health check
	r.GET("/", StatusHandler)

	// Swagger
	r.GET("/swagger/*any", gin.WrapH(httpSwagger.Handler()))

	// User route
	userRepo := userRepository.NewUserRepository()
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)
	user.GetUserRoutes(r, userHandler)

	// 3D Modeling route
	// modelingCollection := db.Collection("modeling3d")
	// modelRepo := &modelingRepository.ModelingRepository{Collection: modelingCollection}
	// modelService := modeling3d.NewModelingService(modelRepo)
	// modelHandler := modeling3d.NewModelingHandler(modelService)

	modelRepo := modelingRepository.NewModelingRepository()
	modelService := modeling3d.NewModelingService(modelRepo)
	modelHandler := modeling3d.NewModelingHandler(modelService)
	modeling3d.GetModelRoutes(r, modelHandler)

	// start the server
	serverAddress := host + ":" + port
	fmt.Printf("Server is running at http://%s\n", serverAddress)
	fmt.Printf("Swagger UI is available at http://%s/swagger\n", serverAddress)

	if err := r.Run(":" + port); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
