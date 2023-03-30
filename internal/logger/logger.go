package logger

import (
	"log"
	"os"
)

type Messages struct{}

func File() *os.File {
	f, err := os.OpenFile("./logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}

	return f
}

func Message() Messages {
	return Messages{}
}
