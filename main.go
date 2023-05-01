package main

import (
	"fmt"
	"math/rand"
)

func main() {

	n := rand.Intn(101)
	println(n)
	fmt.Printf("Please guess a number between 1 and 100 ")

	for {
		var guess int
		fmt.Scanf("%d", &guess)
		if guess == n {
			fmt.Println("\nYou've guessed right!")
			break
		}

		fmt.Print("Your guess is incorrect. ")
		if guess > n {
			fmt.Print("Try again < :")
		} else {
			fmt.Print("Try again > :")
		}
	}
}
