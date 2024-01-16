package main

import (
	"fmt"
	"os"
)

func main() {

	count := len(os.Args) - 1

	fmt.Printf("No of arguments are : %v\n", count)
	fmt.Printf("Path: %v\n", os.Args[0])
}
