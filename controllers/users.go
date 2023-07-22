package controllers

import (
	"backend/models"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	users         []*models.User
	errors        []error
	store         *sessions.CookieStore
	authenticated bool
	user          *models.User
}

func (u *UserController) Users() []*models.User {
	return u.users
}
func (u *UserController) Errors() []error {
	return u.errors
}
func (u *UserController) Authenticated() bool {
	return u.authenticated
}
func (u *UserController) User() *models.User {
	return u.user
}

func (u *UserController) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := u.store.Get(c.Request, "auth-session")
		if id, ok := session.Values["id"].(primitive.ObjectID); ok {
			u.authenticated = true
			user, err := models.User{ID: id}.Read()
			if err != nil {
				u.errors = append(u.errors, err)
				log.Println(err.Error())
			}
			u.user = user
		} else {
			u.authenticated = false
			u.user = nil
		}
		c.Next()
	}
}

func NewUserController() *UserController {
	x := &UserController{}
	users, err := models.User{}.All()
	if err != nil {
		x.errors = append(x.errors, err)
		log.Println(err.Error())
		return nil
	}
	x.users = users
	x.store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	x.authenticated = false
	x.user = nil

	return x
}

// http://localhost:8080/api/users
func (u *UserController) routes(r *gin.RouterGroup) {
	r.POST("/", u.create)
	r.GET("/{id}", u.read)
	// r.GET("/", u.all)
	r.PUT("/{id}", tokens{}.check(), u.update)
	r.DELETE("/{id}", tokens{}.check(), u.delete)
	r.POST("/auth", u.login)
}

// Create
func (u *UserController) create(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response{false, "that data is whack yo", err, nil, http.StatusBadRequest}.send(c)
		return
	}

	if err := user.Create(); err != nil {
		response{false, "i couldn't save that bitch to the database", err, nil, http.StatusBadRequest}.send(c)
		return
	}

	response{true, "User created successfully", nil, user, http.StatusCreated}.send(c)
}

// Read
func (u *UserController) read(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response{false, "your id is bad and you should feel bad", err, nil, http.StatusBadRequest}.send(c)
		return
	}
	user, err := models.User{ID: oid}.Read()
	if err != nil {
		response{false, "i couldn't find that bitch", err, nil, http.StatusInternalServerError}.send(c)
		return
	}

	response{true, "success", nil, user, http.StatusOK}.send(c)
}

//	func (u users) all(c *gin.Context) {
//		// c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
//		users, err := models.User{}.All()
//		if err != nil {
//			response{false, err.Error(), nil, http.StatusInternalServerError}.send(c)
//			return
//		}
//		response{true, "success", users, http.StatusOK}.send(c)
//	}

// Update
func (u UserController) update(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response{false, "that id sucks", err, nil, http.StatusBadRequest}.send(c)
		return
	}
	user, err := models.User{ID: oid}.Read()
	if err != nil {
		response{false, "could not find that user", err, nil, http.StatusInternalServerError}.send(c)
		return
	}

	var userUpdate *models.User
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		response{false, "shitty data", err, nil, http.StatusBadRequest}.send(c)
		return
	}

	if err := user.Update(userUpdate); err != nil {
		response{false, "could not update", err, nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "success", nil, user, http.StatusOK}.send(c)
}

// Delete
func (u UserController) delete(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response{false, "you didn't give me a valid id", err, nil, http.StatusBadRequest}.send(c)
		return
	}
	user, err := models.User{ID: oid}.Read()
	if err != nil {
		response{false, "could not get the user", err, nil, http.StatusInternalServerError}.send(c)
		return
	}

	// make sure the user.ID matches the id in the token
	if err := user.Delete(); err != nil {
		response{false, "could not delete", err, nil, http.StatusInternalServerError}.send(c)
		return
	}
	response{true, "success", nil, user, http.StatusOK}.send(c)
}

// Login
func (u UserController) login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response{
			false,
			"can't bind JSON",
			err,
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
			errors.New("user name not found"),
			nil,
			http.StatusNotFound}.send(c)
		return
	}

	// if name exists check if password is correct
	if !existingUser.PasswordMatches(user.Password) {
		response{
			false,
			"dumb ass",
			errors.New("incorrect password"),
			nil,
			http.StatusUnauthorized}.send(c)
		return
	}

	session, _ := u.store.Get(c.Request, "auth-session")
	session.Values["id"] = existingUser.ID
	session.Values["name"] = existingUser.Name
	session.Options.MaxAge = int(time.Hour.Seconds() * 24)
	if err := session.Save(c.Request, c.Writer); err != nil {
		response{
			false,
			"could not save session",
			err,
			nil,
			http.StatusInternalServerError}.send(c)
		return
	}

	u.authenticated = true
	u.user = existingUser

}
