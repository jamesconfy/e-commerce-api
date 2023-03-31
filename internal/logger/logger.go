package logger

import (
	"fmt"
	"log"
	"os"
)

type Messages struct{}

// Internal server error
func (m Messages) InternalServerError(err error) (str string) {
	str = fmt.Sprintf("Internal server error occurred || Error: %v", err)
	return
}

// Validation error
func (m Messages) ValidationError(req any, err error) (str string) {
	str = fmt.Sprintf("Error when validating request || Request: %s || Error: %v", req, err)
	return
}

func New() *os.File {
	f, err := os.OpenFile("./logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}

	return f
}

// func Message() Messages {
// 	return Messages{}
// }
