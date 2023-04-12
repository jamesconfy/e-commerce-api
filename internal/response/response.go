package response

import (
	se "e-commerce/internal/se"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Status       string `json:"status,omitempty"`
	ResponseCode int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	Error        any    `json:"error,omitempty"` //for errors that occur even if request is successful
	Data         any    `json:"data,omitempty"`
	Extra        any    `json:"extra,omitempty"`
}

type SuccessMessage struct {
	Status       string `json:"status,omitempty" swaggertype:"string" example:"success"`
	ResponseCode int    `json:"code,omitempty" swaggertype:"integer" example:"200"`
	Message      string `json:"message,omitempty" swaggertype:"string" example:"fetched successfully"`
	Data         any    `json:"data,omitempty"`
}

type ErrorMessage struct {
	Status       string `json:"status,omitempty" swaggertype:"string" example:"failure"`
	ResponseCode int    `json:"code,omitempty" swaggertype:"integer" example:"400"`
	Message      string `json:"message,omitempty" swaggertype:"string" example:"error when fetching"`
	Error        any    `json:"error,omitempty"`
}

func NewDecodingError(err error) *Message {
	return &Message{
		ResponseCode: 400,
		Message:      "Bad request",
		Error:        err,
	}
}

func Success(c *gin.Context, message string, data any, extra ...any) {
	msg := &Message{
		Status:       "success",
		ResponseCode: http.StatusOK,
		Message:      message,
		Data:         data,
	}

	c.JSON(http.StatusOK, msg)
}

func Success201(c *gin.Context, message string, data any, extra ...any) {
	msg := &Message{
		Status:       "success",
		ResponseCode: http.StatusCreated,
		Message:      message,
		Data:         data,
	}

	c.JSON(http.StatusOK, msg)
}

func Success202(c *gin.Context, message string) {
	msg := &Message{
		Status:       "success",
		ResponseCode: http.StatusAccepted,
		Message:      message,
		Data:         nil,
	}

	c.JSON(http.StatusOK, msg)
}

func Error(c *gin.Context, sErr se.ServiceError) {
	code := getStatusCodeFromSE(sErr.ErrorType)
	msg := &Message{
		Status:       sErr.ErrorType.String(),
		ResponseCode: code,
		Message:      sErr.Description,
		Error:        sErr.Error,
	}

	c.AbortWithStatusJSON(code, msg)
}

func getStatusCodeFromSE(errorType se.Type) int {
	switch errorType {
	case se.ErrBadRequest:
		return http.StatusBadRequest
	case se.ErrConflict:
		return http.StatusConflict
	case se.ErrNotFound:
		return http.StatusNotFound
	case se.ErrForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
