package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Prints a string from the command line by using 
// the Join function to add a space between each 
// argument.