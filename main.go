package main

import (
	"fmt"
	"os"

	"github.com/Olamilekan-12/go-hello-cli/greet"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go-hello-cli <name> [name...]")
		os.Exit(1)
	}

	for _, name := range os.Args[1:] {
		fmt.Println(greet.Greet(name))
	}
}
