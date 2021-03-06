package main

import (
	"os"
	"fmt"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Prints a string from the command line
// by iterating through the range and adding
// each argument to s with a separater.