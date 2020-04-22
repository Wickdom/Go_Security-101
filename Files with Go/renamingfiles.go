package main

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}

}

func main() {
	oldfile := "hackme.txt"
	newfile := "hackingme.txt"

	err := os.Rename(oldfile, newfile)
	check(err)

}
