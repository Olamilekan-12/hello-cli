package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program: ", os.Args[0])
	fmt.Println("All args: ", os.Args)
	fmt.Println("Number of arguments: ", len(os.Args))
}
