package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dedpidgon/go-web-app/pkg/config"
	"github.com/dedpidgon/go-web-app/pkg/controllers"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// App State
	config.Init()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Database connection
	client, err := connect_to_database()
	if err != nil {
		log.Fatal("Error connecting to mongo...", err.Error())
	}
	defer client.Disconnect(context.TODO())
	gob.Register(primitive.ObjectID{})
	models.Init(client)

	// Sessions
	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   3600 * 24,
		Secure:   config.Production,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	// Routes
	app := gin.Default()

	app.StaticFS("/static", http.Dir("./templates/static/"))

	// middleware
	app.Use(sessions.Sessions("user-session", store))
	app.Use(config.SetClientData())

	cors_config := cors.DefaultConfig()
	cors_config.AllowOrigins = []string{"https://localhost:10000"}
	app.Use(cors.New(cors_config))

	controllers.Handle(app)

	fmt.Println("Server listening on port ", os.Getenv("PORT"))
	app.RunTLS(os.Getenv("PORT"), "cert.pem", "key.pem")
}
