package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// message := fmt.Sprintf(randomFormat())
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// return message, nil
}
func Hellos(names []string) (map[string]string, error) {
	// Create a map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received names, calling the Hello function
	// to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil

}

func randomFormat() string {
	// Return one of the greeting messages selected at random.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]

}
