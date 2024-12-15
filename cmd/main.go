package main

import (
	"fmt"
	"os"

	"github.com/mincong-h/jimi-cli/internal/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
