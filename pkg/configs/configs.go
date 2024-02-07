package configs

import "os"

// Config is a struct that contains the configuration for the application.
// @property {string} MongoURI - The URI of the MongoDB database.
// @property {string} Port - The port that the server will listen on.

type Config struct {
	MongoURL string
	Port     string
}

// It returns a Config struct with the values of the environment variables MONGODB_URL and PORT
func FromEnv() Config {
	config := Config{
		MongoURL: os.Getenv("MONGO_URL"),
		Port:     os.Getenv("PORT"),
	}
	return config
}
