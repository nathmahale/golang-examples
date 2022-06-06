package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin", "Talinki"}

	message, err := greetings.HelloNamed(names)
	if err != nil {
		log.Fatal(err)
	}

	//if no error
	fmt.Println(message)
}
