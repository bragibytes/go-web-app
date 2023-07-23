package controllers

import (
	"net/http"

	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user controller
type user_controller struct {
	authenticated           bool
	authenticated_user_name string
	authenticated_user_id   primitive.ObjectID
	errors                  []string
	ip_address              string
}

func new_user_controller() *user_controller {
	x := &user_controller{}
	x.authenticated = false
	x.authenticated_user_name = "anon"
	x.authenticated_user_id = primitive.NilObjectID

	return x
}

// routes
func (uc *user_controller) use(r *gin.RouterGroup) {
	r.POST("/", uc.create)
	r.GET("/{id}", uc.get_one)
	r.GET("/", uc.get_all)
	r.PUT("/", uc.update)
	r.DELETE("/", uc.delete)
	r.POST("/auth", uc.login)
	r.PUT("/auth", uc.logout)
}

// crud
func (uc *user_controller) create(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response{
			"warning",
			err.Error(),
			nil,
			http.StatusBadRequest,
		}.send(c)
		return
	}

	if err := user.Save(); err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusBadRequest,
		}.send(c)
		return
	}

	response{
		"success",
		"User Created!",
		user,
		http.StatusCreated,
	}.send(c)
}
func (uc *user_controller) get_one(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusBadRequest,
		}.send(c)
		return
	}

	user, err := models.GetOneUser(oid)
	if err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusInternalServerError,
		}.send(c)
		return
	}

	response{
		"success",
		"Got your user!",
		user,
		http.StatusOK,
	}.send(c)
}
func (uc *user_controller) get_all(c *gin.Context) {

	users, err := models.GetAllUsers()
	if err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusInternalServerError,
		}.send(c)
		return
	}
	response{
		"success",
		"Here are your users!",
		users,
		http.StatusOK,
	}.send(c)
}
func (uc *user_controller) update(c *gin.Context) {
	if !uc.authenticated {
		response{
			"warning",
			"You're not authorized to do that",
			nil,
			http.StatusUnauthorized,
		}.send(c)
		return
	}

	user, err := models.GetOneUser(uc.authenticated_user_id)
	if err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusInternalServerError,
		}.send(c)
		return
	}
	var userUpdate *models.User
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		response{
			"warning",
			err.Error(),
			nil,
			http.StatusBadRequest,
		}.send(c)
		return
	}
	if err := user.Update(userUpdate); err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusInternalServerError,
		}.send(c)
		return
	}
	response{
		"success",
		"User updated successfully",
		user,
		http.StatusOK,
	}.send(c)
}
func (uc *user_controller) delete(c *gin.Context) {
	if !uc.authenticated {
		response{
			"warning",
			"You're not authorized to do that",
			nil,
			http.StatusUnauthorized,
		}.send(c)
		return
	}
	user, err := models.GetOneUser(uc.authenticated_user_id)
	if err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusInternalServerError,
		}.send(c)
		return
	}
	if err := user.Delete(); err != nil {
		response{
			"error",
			err.Error(),
			nil,
			http.StatusInternalServerError,
		}.send(c)
		return
	}
	response{
		"success",
		"User deleted successfully",
		user,
		http.StatusOK,
	}.send(c)
}

// authentication
func (uc *user_controller) login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response{
			"warning",
			err.Error(),
			nil,
			http.StatusBadRequest}.send(c)
		return
	}

	// check if username exists
	existingUser := &models.User{Name: user.Name}
	if !existingUser.Exists() {
		response{
			"warning",
			"Username not found",
			nil,
			http.StatusNotFound}.send(c)
		return
	}

	// if name exists check if password is correct
	if !existingUser.PasswordMatches(user.Password) {
		response{
			"warning",
			"Incorrect password",
			nil,
			http.StatusUnauthorized}.send(c)
		return
	}

	ses := sessions.Default(c)
	ses.Set("mongo_id", existingUser.ID)
	ses.Set("username", existingUser.Name)
	if err := ses.Save(); err != nil {
		response{
			"error",
			"Could not save session",
			nil,
			http.StatusInternalServerError,
		}.send(c)
		return
	}
	response{
		"success",
		"Successfully logged in",
		existingUser,
		200,
	}.send(c)
}
func (uc *user_controller) logout(c *gin.Context) {
	ses := sessions.Default(c)
	ses.Set("mongo_id", primitive.NilObjectID)
	ses.Set("username", "")
	ses.Save()
	response{
		"success",
		"im not sure if that worked...",
		nil,
		200,
	}.send(c)
}

// template data
func (uc *user_controller) UserErrors() []string {
	return uc.errors
}
func (uc *user_controller) IsAuthenticated() bool {
	return uc.authenticated
}
func (uc *user_controller) UserName() string {
	return uc.authenticated_user_name
}
func (uc *user_controller) UserID() primitive.ObjectID {
	return uc.authenticated_user_id
}
func (uc *user_controller) GetAllUsers() []*models.User {

	users, err := models.GetAllUsers()
	if err != nil {
		uc.add_error(err)
	}
	return users
}
func (uc *user_controller) GetIP() string {
	return uc.ip_address
}

// helpers
func (uc *user_controller) add_error(err error) {
	uc.errors = append(uc.errors, err.Error())
}

// middleware
func (uc *user_controller) SetUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		remote_ip := c.Request.RemoteAddr
		ses := sessions.Default(c)
		ip := ses.Get("remote_ip")
		if ip == "" {
			// new visitor
			ses.Set("remote_ip", remote_ip)
		}
		uc.ip_address = remote_ip
		mongo_id := ses.Get("mongo_id")
		username := ses.Get("username")
		if mongo_id == nil && username == nil {
			// not logged in
			uc.authenticated = false
		} else {
			// logged in
			uc.authenticated = true
			uc.authenticated_user_name = username.(string)
			uc.authenticated_user_id = mongo_id.(primitive.ObjectID)
		}

	}
}
