package handler

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/service"
	"strconv"

	"e-commerce/internal/response"
	se "e-commerce/internal/se"

	"github.com/gin-gonic/gin"
)

const defaultCookieName = "Authorization"

type UserHandler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	Get(c *gin.Context)
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
	Logout(c *gin.Context)
	ClearAuth(c *gin.Context)
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
// @Router	/users/profile [get]
func (u *userHandler) Get(c *gin.Context) {
	user, err := u.userSrv.GetById(c.GetString("userId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "User gotten successfully", user, 1)
}

// Get User godoc
// @Summary	Get user by id route
// @Description	Get user by id
// @Tags	User
// @Produce	json
// @Param	page	query	int	false	"Page number"
// @Success	200  {object}  response.SuccessMessage{data=[]models.User}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users [get]
func (u *userHandler) GetAll(c *gin.Context) {
	page, _ := c.GetQuery("page")
	if page == "" {
		page = "1"
	}

	pageI, er := strconv.Atoi(page)
	if er != nil {
		err := se.Internal(er, "Error when converting string to integer")
		response.Error(c, *err)
	}

	users, err := u.userSrv.GetAll(pageI)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Users gotten successfully", users, len(users))
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

// Edit User godoc
// @Summary	Edit user route
// @Description	Edit user
// @Tags	User
// @Produce	json
// @Success	200  {object}	response.SuccessMessage{data=models.User}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/profile [patch]
// @Security ApiKeyAuth
func (u *userHandler) Edit(c *gin.Context) {
	var req forms.EditUser

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	userId := c.GetString("userId")

	user, err := u.userSrv.Edit(&req, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "User edited successfully", user)
}

// Delete User godoc
// @Summary	Delete user route
// @Description	Delete user
// @Tags	User
// @Produce	json
// @Success	200  {string}	string	"User deleted successfully"
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/profile [delete]
// @Security ApiKeyAuth
func (u *userHandler) Delete(c *gin.Context) {
	userId := c.GetString("userId")

	err := u.userSrv.Delete(userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	setCookie(c, "", -1)
	response.Success202(c, "User deleted successfully")
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
	err := u.userSrv.DeleteAuth(c.GetString("userId"), getAuth(c))
	if err != nil {
		response.Error(c, *err)
		return
	}

	setCookie(c, "", -1)
	response.Success201(c, "Logged out successfully", nil)
}

// Clear Login Auth godoc
// @Summary	Clear Login Auth route
// @Description	Clear user auth
// @Tags	User
// @Produce	json
// @Success	200  {string}	string	"Logged out from all other device successfully"
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/users/profile/clear [post]
// @Security ApiKeyAuth
func (u *userHandler) ClearAuth(c *gin.Context) {
	err := u.userSrv.ClearAuth(c.GetString("userId"), getAuth(c))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success201(c, "Logged out from all other device successfully", nil)
}

func NewUserHandler(userSrv service.UserService) UserHandler {
	return &userHandler{userSrv: userSrv}
}

// Auxillary function
func setCookie(c *gin.Context, value string, max_age int) {
	c.SetCookie(defaultCookieName, value, 0, "/", "", false, true)
}

func getAuth(c *gin.Context) (auth string) {
	auth, _ = c.Cookie("Authorization")

	if auth != "" {
		return
	}

	auth = c.GetHeader("Authorization")
	return
}
