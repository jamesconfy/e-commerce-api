package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error404() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"description": "Page not found.",
			"status":      http.StatusNotFound,
		})
	}
}
