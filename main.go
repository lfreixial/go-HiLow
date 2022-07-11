package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	lives := 3
	score := 0
	var answer string

	for lives != 0 {
		rand.Seed(time.Now().UnixNano())
		rng := rand.Intn(20)
		rngNext := rand.Intn(20)
		fmt.Printf("Lives: %v | Score: %v\n", lives, score)
		fmt.Printf("Is the next number higher or lower than %v\n", rng)
		fmt.Print("1: Higher\n2: Lower\n-> ")
		fmt.Scanln(&answer)
		choice, err := strconv.ParseInt(answer, 5, 12)
		if err != nil || choice <= 0 || choice > 2 {
			for {
				fmt.Println("Error - Choice not valid. Please select one from the following")
				fmt.Print("1: Higher\n2: Lower\n-> ")
				fmt.Scanln(&answer)
				choice, err = strconv.ParseInt(answer, 5, 12)
				if err == nil && choice > 0 && choice <= 2 {
					break
				}
				_ = choice
			}

		}
		if choice == 1 && rngNext >= rng {
			score += 1
		} else if choice == 2 && rngNext <= rng {
			score += 1
		} else {
			lives -= 1
		}

	}

	fmt.Printf("You scored %v\n", score)
}
