package handler

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/service"

	"e-commerce/internal/response"
	se "e-commerce/internal/se"

	"github.com/gin-gonic/gin"
)

const defaultCookieName = "Authorization"

type UserHandler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	GetById(c *gin.Context)
	Logout(c *gin.Context)
	// ResetPassword(c *gin.Context)
	// ValidateToken(c *gin.Context)
	// ChangePassword(c *gin.Context)
}

type userHandler struct {
	userSrv service.UserService
}

// Register User godoc
// @Summary	Register route
// @Description	Register route
// @Tags	User
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Signup	true "Signup Details"
// @Success	200  {object}  response.SuccessMessage
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/signup [post]
func (u *userHandler) Create(c *gin.Context) {
	var req forms.Signup

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	user, err := u.userSrv.Add(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "User created successfully", user)
}

// Login User godoc
// @Summary	Login route
// @Description	Login route
// @Tags	User
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Login	true "Login Details"
// @Success	200  {object}  response.SuccessMessage
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/login [post]
func (u *userHandler) Login(c *gin.Context) {
	var req forms.Login

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	auth, err := u.userSrv.Login(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	setCookie(c, auth.AccessToken, 0)
	response.Success(c, "User logged in successfully", auth)
}

// Get User godoc
// @Summary	Get user by id route
// @Description	Get user by id
// @Tags	User
// @Produce	json
// @Param	userId	path	string	true	"User id"
// @Success	200  {object}  response.SuccessMessage{data=models.User}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/:userId [get]
func (u *userHandler) GetById(c *gin.Context) {
	user, err := u.userSrv.GetById(c.Param("userId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "User gotten successfully", user, 1)
}

// Logout User godoc
// @Summary	Logout user route
// @Description	Logout user
// @Tags	User
// @Produce	json
// @Success	200  {string}	string	"Logged out successfully"
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/logout [post]
// @Security ApiKeyAuth
func (u *userHandler) Logout(c *gin.Context) {
	userId := c.GetString("userId")
	err := u.userSrv.DeleteToken(userId)
	if err != nil {
		response.Error(c, *err)
	}

	setCookie(c, "", -1)
	response.Success201(c, "Logged out successfully", nil)
}

func NewUserHandler(userSrv service.UserService) UserHandler {
	return &userHandler{userSrv: userSrv}
}

// Auxillary function
func setCookie(c *gin.Context, value string, max_age int) {
	c.SetCookie(defaultCookieName, value, 0, "/", "", false, true)
}
