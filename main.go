package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go-hello-cli <name> [name...]")
		os.Exit(1)
	}

	for _, name := range os.Args[1:] {
		fmt.Println(greet(name))
	}
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
