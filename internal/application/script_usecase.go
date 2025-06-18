package application

import (
	"fmt"
	"time"
)

type CommandHandler struct{}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

func (h *CommandHandler) Check() {
	timestamp := time.Now().Format("2006-01-02")
	fmt.Printf("Prueba iniciada a las %s", timestamp)

}
