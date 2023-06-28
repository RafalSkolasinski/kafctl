package main

import (
	"log"

	"github.com/rafalskolasinski/kctl/cmd"
)

func main() {
	log.SetFlags(0)
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
