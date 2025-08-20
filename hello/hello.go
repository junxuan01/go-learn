package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	// 引入 greetings 包
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request a greeting message.
	message, err := greetings.Hellos(names)
	// if An error was returned, print it to the console and exit.
	if err != nil {
		fmt.Printf("error: %v\n", err)
		log.Fatal(err)
	}
	fmt.Println(message)
}
