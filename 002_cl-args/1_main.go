// Prints command-line arguements
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Prints a string from the command line by
// taking each argument from the list and appending
// it to s with a separator.
