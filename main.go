package main

import (
	"backend/controllers"
	"backend/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend/views"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port = ":8000"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	client, err := connectToDB()
	if err != nil {
		log.Fatal("Error connecting to mongo...", err.Error())
	}
	defer client.Disconnect(context.TODO())

	models.Init(client)

	router := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:4200"}
	// config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// router.Use(cors.New(config))
	controllers.Init(router.Group("/api/v1/"))

	router.StaticFS("/static", http.Dir("./views/static/"))

	views.Init(router)

	fmt.Println("Server is running on port ", port)
	router.Run(port)
}

func connectToDB() (*mongo.Client, error) {
	// Replace the connection string with your MongoDB Atlas connection string
	connectionString := os.Getenv("MONGO_URI")

	// Set up client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check if the connection was successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
