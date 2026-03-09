package main

import (
	"log"
	"github.com/adr/ad-guidance-tool/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
