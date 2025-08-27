package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin" //gin framework
	"github.com/joho/godotenv" //local variable on .env files
	httpSwagger "github.com/swaggo/http-swagger" //auto-API docs writter
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

	fmt.Println("CLIENT_ID:", os.Getenv("CLIENT_ID"))
	fmt.Println("CLIENT_SECRET:", os.Getenv("CLIENT_SECRET"))

	dbConfig.InitMongo()

	r := gin.Default()
	r.Static("/static", "./")

	r.Use(func(c *gin.Context){
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

	// userRepo := userRepository.NewUserRepository()
	// situationRepo := situationRepository.NewSituationRepository()
	// favoriteService := favorite.NewFavoriteService(userRepo, situationRepo)
	// favoriteHandler := favorite.NewFavoriteHandler(favoriteService)
	// favorite.RegisterFavoriteRoutes(r, favoriteHandler)

	// start the server
	serverAddress := host + ":" + port
	fmt.Printf("Server is running at http://%s\n", serverAddress)
	fmt.Printf("Swagger UI is available at http://%s/swagger\n", serverAddress)

	if err := r.Run(":" + port); err != nil {
		fmt.Println("Error starting server:", err)
	}

}