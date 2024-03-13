package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// rudimentary error handling call
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// returns a random string using the rand exported name from math
func randomFormat() string {
	formats := []string{
		"Hey, %v! Welcome!",
		"Great to see you, %v!",
		"Ahoy, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
