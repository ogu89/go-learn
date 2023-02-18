package gretings

import "fmt"

// Hello returns a greetings for the named person
func Hello(name string) string{
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

