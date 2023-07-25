package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
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
		response.BadReq(c, err)
		return
	}

	if validation_errors := user.Validate(); len(validation_errors) > 0 {
		response.ValidationErrors(c, validation_errors)
		return
	}

	if err := user.Save(); err != nil {
		response.ServerErr(c, err)
		return
	}

	response.Created(c, "user", user)
}
func (uc *user_controller) get_one(c *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		response.BadReq(c, err)
		return
	}

	user, err := models.GetOneUser(oid)
	if err != nil {
		response.NotFound(c, "user", err.Error())
		return
	}

	response.OK(c, "got your user", user)
}
func (uc *user_controller) get_all(c *gin.Context) {

	users, err := models.GetAllUsers()
	if err != nil {
		response.NotFound(c, "users", err.Error())
		return
	}
	response.OK(c, "here are your users", users)
}
func (uc *user_controller) update(c *gin.Context) {
	if !uc.authenticated {
		response.Unauthorized(c, "i cant let you do that")
		return
	}

	user, err := models.GetOneUser(uc.authenticated_user_id)
	if err != nil {
		response.ServerErr(c, err)
		return
	}
	var userUpdate *models.User
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		response.BadReq(c, err)
		return
	}
	if err := user.Update(userUpdate); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "successfully updated user", user)
}
func (uc *user_controller) delete(c *gin.Context) {
	if !uc.authenticated {
		response.Unauthorized(c, "i cant let you do that")
		return
	}
	userToDelete, err := models.GetOneUser(uc.authenticated_user_id)
	if err != nil {
		response.ServerErr(c, err)
		return
	}
	if err := userToDelete.Delete(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "successfully deleted user", userToDelete)
}

// authentication
func (uc *user_controller) login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadReq(c, err)
		return
	}

	// check if username exists
	existingUser := &models.User{Name: user.Name}
	if !existingUser.Exists() {
		response.NotFound(c, "user", "error in login handler")
		return
	}

	// if name exists check if password is correct
	if !existingUser.PasswordMatches(user.Password) {
		response.Unauthorized(c, "wrong password")
		return
	}

	ses := sessions.Default(c)
	ses.Set("mongo_id", existingUser.ID)
	ses.Set("username", existingUser.Name)
	if err := ses.Save(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "logged in", existingUser)
}
func (uc *user_controller) logout(c *gin.Context) {
	ses := sessions.Default(c)
	ses.Clear()
	if err := ses.Save(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "logged out", nil)
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
