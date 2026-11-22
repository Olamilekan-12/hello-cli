package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go-hello-cli <name>")
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Printf("Hello, %s!\n", name)
}
