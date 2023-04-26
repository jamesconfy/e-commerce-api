package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorHandler interface {
	Error404(c *gin.Context)
	MethodNotAllowed(c *gin.Context)
}

type errorHandler struct {
}

func (e *errorHandler) Error404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"description": "Page not found.", "status": http.StatusNotFound})
}

func (e *errorHandler) MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"description": "Method not allowed", "status": http.StatusMethodNotAllowed})
}

func NewErrorHandler() ErrorHandler {
	return &errorHandler{}
}
