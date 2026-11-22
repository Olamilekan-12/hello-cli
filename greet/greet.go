package greet

import "fmt"

// Greet returns greeting for the given name
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
