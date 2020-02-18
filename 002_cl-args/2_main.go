package main

import (
	"fmt"
	"os"
)

func main() {

	argsOnly := os.Args[0]
	argsWithProg := os.Args
	argsApple := os.Args[:len(os.Args)]
	argsBahama := os.Args[1:]
	argsCat := os.Args[1:len(os.Args)] 
	argsDinner := os.Args[1:5]
	argsElephant := os.Args[1:4]
	argsFrance := os.Args[2:]
	argsGravy := os.Args[:2]
	argsHappy := os.Args[2]
	argsIgloo := os.Args[2:4]
	

	fmt.Println(argsOnly)
	fmt.Println(argsWithProg)
	fmt.Println(argsApple)
	fmt.Println(argsBahama)
	fmt.Println(argsCat)
	fmt.Println(argsDinner)
	fmt.Println(argsElephant)
	fmt.Println(argsFrance)
	fmt.Println(argsGravy)
	fmt.Println(argsHappy)
	fmt.Println(argsIgloo)
}

//	Output from binary "this is a test" 

//	./2_main
//	[./2_main this is a test]
//	[./2_main this is a test]
//	[this is a test]
//	[this is a test]
//	[this is a test]
//	[this is a]
//	[is a test]
//	[./2_main this]
//	is
//	[is a]


