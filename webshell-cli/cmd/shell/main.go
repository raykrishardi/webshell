package main

import (
	"log"

	"github.com/raykrishardi/webshell-cli/internal/pkg/shell/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}
