package main

import (
	"log"

	cmd "github.com/hellozimi/cenv/cmd/cenv"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
