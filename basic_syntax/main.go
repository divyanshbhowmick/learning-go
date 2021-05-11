package basic_syntax

import "fmt"

func main() {
	fmt.Println("Hello World! Let's GO! 😎")
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hey, %v! Welcome! 😎", name)
	return message
}
