package handlers

import (
	"github.com/dedpidgon/go-web-app/models"
	"github.com/gorilla/sessions"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userState struct {
	Authenticated bool
	User          *models.User
	store         *sessions.CookieStore
}

func newUserState() *userState {
	x := &userState{}
	x.Authenticated = false
	x.User = nil
	x.store = sessions.NewCookieStore()

	return x
}

// http://localhost:8080/api/users
func user_crud(r *gin.RouterGroup) {
	r.POST("/", create_user)
	r.GET("/{id}", read_user)
	r.GET("/", read_all_users)
	r.PUT("/{id}", update_user)
	r.DELETE("/{id}", destroy_user)
	r.POST("/auth", login_user)
}

// Create
func create_user(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response{false, "that data is whack yo", err.Error(), nil, http.StatusBadRequest}.send(c)
		return
	}

	if err := user.Create(); err != nil {
		response{false, "i couldn't save that bitch to the database", err.Error(), nil, http.StatusBadRequest}.send(c)
		return
	}

	response{true, "User created successfully", "", user, http.StatusCreated}.send(c)
}

// Read
func read_user(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response{false, "your id is bad and you should feel bad", err.Error(), nil, http.StatusBadRequest}.send(c)
		return
	}
	user, err := models.User{ID: oid}.Read()
	if err != nil {
		response{false, "i couldn't find that bitch", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}

	response{true, "success", "", user, http.StatusOK}.send(c)
}

func read_all_users(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
	users, err := models.User{}.All()
	if err != nil {
		response{false, "", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "success", "", users, http.StatusOK}.send(c)
}

// Update
func update_user(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response{false, "that id sucks", err.Error(), nil, http.StatusBadRequest}.send(c)
		return
	}
	user, err := models.User{ID: oid}.Read()
	if err != nil {
		response{false, "could not find that user", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}

	var userUpdate *models.User
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		response{false, "shitty data", err.Error(), nil, http.StatusBadRequest}.send(c)
		return
	}

	if err := user.Update(userUpdate); err != nil {
		response{false, "could not update", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "success", "", user, http.StatusOK}.send(c)
}

// Delete
func destroy_user(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response{false, "you didn't give me a valid id", err.Error(), nil, http.StatusBadRequest}.send(c)
		return
	}
	user, err := models.User{ID: oid}.Read()
	if err != nil {
		response{false, "could not get the user", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}

	// make sure the user.ID matches the id in the token
	if err := user.Delete(); err != nil {
		response{false, "could not delete", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "success", "", user, http.StatusOK}.send(c)
}

// Login
func login_user(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response{
			false,
			"can't bind JSON",
			err.Error(),
			nil,
			http.StatusBadRequest}.send(c)
		return
	}

	// check if user name exists
	existingUser := models.User{
		Name: user.Name,
	}.Exists()
	if existingUser == nil {
		response{
			false,
			"user not found",
			"user name not found",
			nil,
			http.StatusNotFound}.send(c)
		return
	}

	// if name exists check if password is correct
	if !existingUser.PasswordMatches(user.Password) {
		response{
			false,
			"dumb ass",
			"incorrect password",
			nil,
			http.StatusUnauthorized}.send(c)
		return
	}

}
