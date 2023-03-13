package userHandler

import (
	"log"
	"net/http"

	"e-commerce/internal/models/responseModels"
	"e-commerce/internal/models/userModels"
	"e-commerce/internal/service/userService"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSrv userService.UserService
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var req userModels.CreateUserReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	user, errU := h.userSrv.CreateUser(&req)
	if errU != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error creating user", errU, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "User created successfully", user, nil))
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var req userModels.LoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	user, errU := h.userSrv.Login(&req)
	if errU != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error logging user", errU, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "User logged in successfully", user, nil))
}

func (h *userHandler) ResetPassword(c *gin.Context) {
	var req userModels.ResetPasswordReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	token, errP := h.userSrv.ResetPassword(&req)
	if errP != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error resetting password", errP, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Check your email for reset token", token, nil))
}

func (h *userHandler) ValidateToken(c *gin.Context) {
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
	// 	return
	// }

	userId := string(c.Query("user_id"))
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "No user id provided in path", nil, nil))
		return
	}

	tokenId := string(c.Query("token"))
	if tokenId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "No token provided in path", nil, nil))
		return
	}

	token, errT := h.userSrv.ValidateToken(userId, tokenId) //, tokenId)
	if errT != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when validating token", errT, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Token validated successfully", token, nil))
}

func (h *userHandler) ChangePassword(c *gin.Context) {
	var req userModels.ChangePasswordReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err.Error(), nil))
		return

	}

	userId := string(c.Query("user_id"))
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "No user id provided in path", nil, nil))
		return
	}

	errT := h.userSrv.ChangePassword(userId, &req) //, tokenId)
	// if errT.Description == "Bad Input Request" {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Passwords do not match", errT, nil))
	// 	return
	// }

	if errT != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when changing password", errT, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Password changed successfully", nil, nil))
}

func NewUserHandler(userSrv userService.UserService) *userHandler {
	return &userHandler{userSrv: userSrv}
}
