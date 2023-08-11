package main

import (
	"log"

	"github.com/rafalskolasinski/kafctl/cmd"
)

func main() {
	log.SetFlags(0)
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
