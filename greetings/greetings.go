package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome! ",
		"Wonderful, that you graced us with your presence %v",
		"Howdy %v...",
		"Konnichiwa %v",
		"Its great to have you %v San",
	}

	return formats[rand.Intn(len(formats))]
}
