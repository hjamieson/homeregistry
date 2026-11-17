package main

import (
	"github.com/hjamieson/homeregistry/internal/platform"
	"log"
)

func main() {
	log.Print("lookup machineID from OS...")
	machineID, err := platform.MachineID()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("machineID found: %q", machineID)
}
