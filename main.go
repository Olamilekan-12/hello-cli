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
		fmt.Printf("Hello, %s!\n", name)
	}
}
