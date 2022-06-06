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

func HelloNamed(names []string) (map[string]string, error) {
	//maps for names with messages
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome! \n",
		"Good to have you %v \n",
		"Howdy %v...\n",
		"Konnichiwa %v\n",
		"Its great to have you %v San\n",
	}

	return formats[rand.Intn(len(formats))]
}
