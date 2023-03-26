package handler

import (
	"net/http"

	"e-commerce/internal/service/homeService"

	"github.com/gin-gonic/gin"
)

type homeHandler struct {
	homeSrv homeService.HomeService
}

func (h *homeHandler) Home(c *gin.Context) {
	message, err := h.homeSrv.CreateHome()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Internal error"})
	}

	c.JSON(http.StatusOK, gin.H{"msg": message})
}

func NewHomeHandler(homeSrv homeService.HomeService) *homeHandler {
	return &homeHandler{homeSrv: homeSrv}
}
