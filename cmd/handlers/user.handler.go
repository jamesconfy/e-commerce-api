package handler

import (
	"errors"

	"e-commerce/internal/forms"
	"e-commerce/internal/service"

	"e-commerce/internal/response"
	se "e-commerce/internal/serviceerror"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	GetById(c *gin.Context)
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
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	request	body	userModels.CreateUserReq	true "Signup Details"
// @Success	200  {object}  userModels.CreateUserRes
// @Failure	400  {object}  errorModels.ServiceError
// @Failure	404  {object}  errorModels.ServiceError
// @Failure	500  {object}  errorModels.ServiceError
// @Router	/users [post]
func (u *userHandler) Create(c *gin.Context) {
	var req forms.Signup

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, *se.Validating(errors.New("Details not found")))
		return
	}

	// if err := u.userSrv.Validate(req); err != nil {
	// 	response.Error(c, *se.NewValidating(err))
	// 	return
	// }

	user, err := u.userSrv.Create(&req)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success("User created successfully", user)
}

// Login User godoc
// @Summary	Login route
// @Description	Login route
// @Tags	Users
// @Accept	json
// @Produce	json
// @Param	request	body	userModels.LoginReq	true "Login Details"
// @Success	200  {object}  userModels.LoginRes
// @Failure	400  {object}  errorModels.ServiceError
// @Failure	404  {object}  errorModels.ServiceError
// @Failure	500  {object}  errorModels.ServiceError
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

	response.Success("User logged in successfully", auth, nil)
}

func (u *userHandler) GetById(c *gin.Context) {
	user, err := u.userSrv.GetById(c.Param("userId"))
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success("User gotten successfully", user, nil)
}

// func (h *userHandler) ResetPassword(c *gin.Context) {
// 	var req userModels.ResetPasswordReq

// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
// 		return
// 	}

// 	token, errP := h.userSrv.ResetPassword(&req)
// 	if errP != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error resetting password", errP, nil))
// 		return
// 	}

// 	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Check your email for reset token", token, nil))
// }

// func (h *userHandler) ValidateToken(c *gin.Context) {
// 	// err := c.ShouldBindJSON(&req)
// 	// if err != nil {
// 	// 	c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
// 	// 	return
// 	// }

// 	userId := string(c.Query("user_id"))
// 	if userId == "" {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "No user id provided in path", nil, nil))
// 		return
// 	}

// 	tokenId := string(c.Query("token"))
// 	if tokenId == "" {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "No token provided in path", nil, nil))
// 		return
// 	}

// 	token, errT := h.userSrv.ValidateToken(userId, tokenId) //, tokenId)
// 	if errT != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when validating token", errT, nil))
// 		return
// 	}

// 	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Token validated successfully", token, nil))
// }

// func (h *userHandler) ChangePassword(c *gin.Context) {
// 	var req userModels.ChangePasswordReq

// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		log.Println(err.Error())
// 		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err.Error(), nil))
// 		return

// 	}

// 	userId := string(c.Query("user_id"))
// 	if userId == "" {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "No user id provided in path", nil, nil))
// 		return
// 	}

// 	errT := h.userSrv.ChangePassword(userId, &req) //, tokenId)
// 	// if errT.Description == "Bad Input Request" {
// 	// 	c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Passwords do not match", errT, nil))
// 	// 	return
// 	// }

// 	if errT != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when changing password", errT, nil))
// 		return
// 	}

// 	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Password changed successfully", nil, nil))
// }

func NewUserHandler(userSrv service.UserService) UserHandler {
	return &userHandler{userSrv: userSrv}
}
