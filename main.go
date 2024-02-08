package main

import (
	"conversationserver/api/routes"
	"os"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// `app := gin.Default()` is creating a new instance of the Gin framework's default router.
	app := gin.Default()

	// The code `app.Use(cors.New(cors.Config{...}))` is configuring Cross-Origin Resource Sharing (CORS)
	// middleware for the Gin framework.
	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowHeaders:     []string{"Accept, Content-Type, Content-Length, Token"},
				AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			}))

	// `godotenv.Load()` It is used to load environment variables from a `.env` file into the current environment.
	godotenv.Load()

	routes.CreateChatRoomRoute(app)

	app.Run("localhost:" + os.Getenv("PORT"))
}
