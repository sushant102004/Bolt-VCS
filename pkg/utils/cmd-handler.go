/*
	PoF: - Provide a blueprint that accept a command and execute given function.
*/

package utils

import (
	"log"
)

type CMDHandler struct {
	commandsList map[string]func()
	l            *log.Logger
}

func NewCMDHandler() *CMDHandler {
	return &CMDHandler{
		commandsList: make(map[string]func()),
		l:            &log.Logger{},
	}
}

func (h *CMDHandler) AddCommand(cmd string, function func()) {
	h.commandsList[cmd] = function
}

func (h *CMDHandler) Listen(cmd string) {
	for key, value := range h.commandsList {
		if key == cmd {
			value()
		}
	}
}
