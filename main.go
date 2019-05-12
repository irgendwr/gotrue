package main

import (
	"log"

	"github.com/irgendwr/gotrue/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
