package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name!!")
	}

	// if a name was received
	message := fmt.Sprintf("Hi, %v. Welkom!", name)
	return message, nil
}
