package main

import (
	"context"
	"conversationserver/pkg/configs"
	"log"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// The line `config := configs.FromEnv()` is calling the `FromEnv()` function from the `configs` package and assigning the returned value to the `config` variable.
	config := configs.FromEnv()

	// The line `client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURL))` is establishing a connection to a MongoDB database.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURL))

	// The code `if err != nil { log.Panic(err) }` is checking if there was an error during the connection to the MongoDB database.
	if err != nil {
		log.Panic(err)
	}

	// The line `db := client.Database("conversation_server")` is creating a handle to the "conversation_server" database in MongoDB.
	db := client.Database("conversation_server")

	println(db)
	app.Run("http://localhost:" + config.Port)
}
