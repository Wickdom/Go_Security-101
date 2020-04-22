package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	info, err := os.Stat("hackme.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("File Permissions: ", info.Mode())
	fmt.Println("Last Modified:", info.ModTime())
	fmt.Println("Directory: ", info.IsDir())
	fmt.Printf("System interface type: %T/n", info.Sys())
	fmt.Printf("System info: %+v\n\n", info.Sys())
}
