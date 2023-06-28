package main

import (
	"log"

	"github.com/rafalskolasinski/kctl-golang/cmd/kctl"
)

func main() {
	log.SetFlags(0)
	err := kctl.Execute()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
