package main

import (
	"context"
	"fmt"
	"github.com/dedpidgon/go-web-app/pkg/controllers"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
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
	controllers.InitData()
	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   3600 * 24,              // Set the session age to 1 hour (in seconds)
		Secure:   controllers.Production, // Set this to true in production to use secure (HTTPS) connections
		HttpOnly: true,                   // This prevents JavaScript from accessing the cookie
		SameSite: http.SameSiteLaxMode,
	})
	app := gin.Default()

	app.StaticFS("/static", http.Dir("./templates/static/"))

	// middleware
	app.Use(sessions.Sessions("user-session", store))
	app.Use(controllers.UserController.SetUserData())

	controllers.InitRoutes(app)
	fmt.Println("Server listining on port ", os.Getenv("PORT"))
	app.Run(os.Getenv("PORT"))
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
