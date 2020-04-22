package main

import (
	"log"
	"os"
)

func main() {
	err := os.Truncate("hackme.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
