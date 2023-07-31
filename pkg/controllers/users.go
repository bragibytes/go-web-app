package controllers

import (
	"fmt"
	"net/url"

	"github.com/dedpidgon/go-web-app/pkg/config"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user controller
type user_controller struct{}

// routes
func (uc *user_controller) use(r *gin.RouterGroup) {
	r.POST("/", uc.create)
	r.GET("/{id}", uc.get_one)
	r.GET("/", uc.get_all)
	r.PUT("/", uc.update)
	r.PUT("/bio", uc.update_bio)
	r.DELETE("/", uc.delete)
	r.POST("/auth", uc.login)
	r.PUT("/auth", uc.logout)
	r.GET("/verify/:token", uc.verify)
	r.GET("/resend/:id", uc.resend_verification)
}

// crud
func (uc *user_controller) create(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadReq(c, err)
		return
	}

	if validation_errors := user.Valid(); len(validation_errors) > 0 {
		response.ValidationErrors(c, validation_errors)
		return
	}

	if err := user.Save(); err != nil {
		response.ServerErr(c, err)
		return
	}

	if err := save_session(c, user.ID, user.Name); err != nil {
		response.ServerErr(c, err)
		return
	}
	if err := email_control.verify(user); err != nil {
		response.Send(c, "info", err.Error()+" verification email not sent", user, 201)
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
	if !config.Client.Authenticated {
		response.Unauthorized(c, "i cant let you do that")
		return
	}

	user, err := models.GetOneUser(config.Client.ID)
	if err != nil {
		response.ServerErr(c, err)
		return
	}
	var user_update *models.User
	if err := c.ShouldBindJSON(&user_update); err != nil {
		response.BadReq(c, err)
		return
	}
	user_update.UpdateFriend(user)
	if validation_errors := user.Valid(); len(validation_errors) > 0 {
		for _, e := range validation_errors {
			fmt.Println("==========" + e + "===========")
		}
		response.ValidationErrors(c, validation_errors)
		return
	}

	if err := user.Update(); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "successfully updated user", user)
}
func (uc *user_controller) delete(c *gin.Context) {
	if !config.Client.Authenticated {
		response.Unauthorized(c, "i cant let you do that")
		return
	}
	userToDelete, err := models.GetOneUser(config.Client.ID)
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

	if err := save_session(c, existingUser.ID, existingUser.Name); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "logged in", nil)
}
func (uc *user_controller) logout(c *gin.Context) {

	if err := delete_session(c); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "logged out", nil)
}

func (uc *user_controller) verify(c *gin.Context) {
	tokenString, err := url.QueryUnescape(c.Param("token"))
	if err != nil {
		response.BadReq(c, err)
		return
	}
	idHex, err := token_control.get_user_id(tokenString)
	if err != nil {
		view_control.bad_token(c)
		fmt.Println(err.Error())
		return
	}
	oid, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		response.BadReq(c, err)
		return
	}
	user, err := models.GetOneUser(oid)
	if err != nil {
		view_control.bad_token(c)
		fmt.Println(err.Error())
		return
	}
	user.Verify()

	// log user in
	if err := save_session(c, user.ID, user.Name); err != nil {
		view_control.bad_token(c)
		fmt.Println(err.Error())
		return
	}
	// send to profile
	view_control.profile(c)
}

func (uc *user_controller) resend_verification(c *gin.Context) {
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
	if err := email_control.verify(user); err != nil {
		response.ServerErr(c, err)
		return
	}
	response.OK(c, "check your email", nil)
}

func (uc *user_controller) update_bio(c *gin.Context) {
	var user_with_bio *models.User
	if err := c.ShouldBindJSON(&user_with_bio); err != nil {
		response.BadReq(c, err)
		return
	}
	user, err := models.GetOneUser(config.Client.ID)
	if err != nil {
		response.NotFound(c, "user", err.Error())
		return
	}
	if err := user.UpdateBio(user_with_bio.Bio); err != nil {
		response.ServerErr(c, err)
		return
	}

	response.OK(c, user.OK, user)
}
