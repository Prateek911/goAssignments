package main

import (
	"fmt"
	"os"
	"strconv"
)

func isEven(n int) (bool, error) {
	if n%2 == 0 {
		return true, nil
	}
	return false, fmt.Errorf("%d is not an even number", n)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}
	input := os.Args[1]
	if num, err := strconv.Atoi(input); err != nil {
		fmt.Println("Not a number")
		os.Exit(1)
	} else if result, err := isEven(num); err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	} else if result {
		fmt.Printf("%d is an even number.\n", num)
	} else {
		fmt.Printf("%d is an odd number.\n", num)
	}
}
