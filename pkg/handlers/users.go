package handlers

import (
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/gorilla/sessions"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ustate *user_state

type user_state struct {
	Authenticated bool
	User          *models.User
	store         *sessions.CookieStore
	Errors        []string
}

func new_user_state() *user_state {
	x := &user_state{}
	x.Authenticated = false
	x.User = nil
	x.store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	x.store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24, // Cookie expiration time in seconds (1 hour)
		HttpOnly: true,
		Secure:   false, // Set this to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
	}

	return x
}

func (us *user_state) GetAllUsers() []*models.User {
	x := &models.User{}
	users, err := x.GetAll()
	if err != nil {
		us.add_error(err.Error())
		return nil
	}
	return users
}
func (us *user_state) soft_cookie_check(c *gin.Context) {

	session, err := us.store.Get(c.Request, "auth-session")
	if err != nil {
		// Session doesn't exist or is a new session, handle it as needed
		us.add_error(err.Error())
		us.clear_user_data()
	}
	id, ok := session.Values["id"]
	if !ok {
		us.add_error("could not find an id in the cookie")
		us.clear_user_data()
	}
	x := &models.User{ID: id.(primitive.ObjectID)}
	// Continue with the request
	us.Authenticated = true
	err = x.PopulateByID()
	if err != nil {
		us.add_error(err.Error())
		us.clear_user_data()
	}
	us.User = x
	c.Next()
}
func (us *user_state) clear_user_data() {
	us.Authenticated = false
	us.User = nil
}
func (us *user_state) add_error(err string) {
	us.Errors = append(us.Errors, err)
}
func (us *user_state) hard_cookie_check(c *gin.Context) {
	session, err := us.store.Get(c.Request, "auth-session")
	if err != nil {
		// Session doesn't exist or is a new session, handle it as needed
		us.clear_user_data()
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	id, ok := session.Values["id"]
	if !ok {
		us.clear_user_data()
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	x := &models.User{ID: id.(primitive.ObjectID)}
	// Continue with the request
	us.Authenticated = true
	if err = x.PopulateByID(); err != nil {
		us.clear_user_data()
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	us.User = x
	c.Next()
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
	user := &models.User{ID: oid}
	if err = user.PopulateByID(); err != nil {
		response{false, "i couldn't find that bitch", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}

	response{true, "success", "", user, http.StatusOK}.send(c)
}

func read_all_users(c *gin.Context) {
	x := &models.User{}
	users, err := x.GetAll()
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
	user := &models.User{ID: oid}
	if err = user.PopulateByID(); err != nil {
		response{false, "could not find that user", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}

	var userUpdate *models.User
	if err = c.ShouldBindJSON(&userUpdate); err != nil {
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
	user := &models.User{ID: oid}
	if err = user.PopulateByID(); err != nil {
		response{false, "could not get the user", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}

	// make sure the user.ID matches the id in the token
	if err = user.Delete(); err != nil {
		response{false, "could not delete", err.Error(), nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "here is the corps of your deleted user", "no errors, but shame on you", user, http.StatusOK}.send(c)
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

	// check if username exists

	existingUser := &models.User{Name: user.Name}

	if !existingUser.Exists() {
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
