package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 10
	usage    = `Program will pick %d random numbers. Your mission is to guess one of those. Wanna play?
	`
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	if guess < 0 {
		fmt.Println("Pleae enter positive number")
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			fmt.Println("YOU WIN!!")
			return
		}
	}
	fmt.Println("You lost try again!!")
}
