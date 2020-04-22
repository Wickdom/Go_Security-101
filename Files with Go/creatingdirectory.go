package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("Main_hack", 0755)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Created the File was", Name())
	/*
		fmt.Println("Created the File was", dir.Name())
		fmt.Println("Directory: ", dir.IsDir())
		fmt.Printf("System interface type: %T/n", dir.Sys())
		fmt.Printf("System info: %+v\n\n", dir.Sys())

	*/
}
