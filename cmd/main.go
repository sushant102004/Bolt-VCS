package main

import (
	"os"

	internal "github.com/sushant102004/bolt-vcs/internal/commands"
	"github.com/sushant102004/bolt-vcs/pkg/utils"
)

func main() {
	cmdHandler := utils.NewCMDHandler()

	cmdHandler.AddCommand("init", internal.Initialize)

	cmdHandler.Listen(os.Args[1])
}
