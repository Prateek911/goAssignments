package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const list = "" + "quick brown fox jumps over the lazy dog"

func main() {

	words := strings.Fields(list)
	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Please enter at least one string")
		return
	}

	for _, q := range query {
		if isString(q) {
			for i, w := range words {
				if q == w {
					fmt.Printf("#%-2d: %q\n", i+1, w)
				}
			}
		} else {
			fmt.Println("Use only Strings!!")
			return
		}
	}

}

func isString(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		_, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return true
		}
	}
	return false
}
