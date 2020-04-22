package main

import (
	"fmt"
	"log"
	"os" //because dealing with the os level this package is imported
)

func main() {
	//Creating the a file and using the log to verify
	//filename is hackme.txt
	File, err := os.Create("hackme.txt")
	//a conditional statement to check
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(File)
	log.Println(File)
	File.Close()

}
