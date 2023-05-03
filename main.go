package main

import (
	"fmt"
	"math/rand"
)

//"math/rand"

func main() {

	guessNumber()  // Gess a number between 1 and 100
	ifSample()     // Demonstrating the use of if, else and else if
	switchSample() // Switch without switch statement
	logicSample()  // Logical Operators
}

// *********** gess a number block ***********
func guessNumber() {

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

// ***********  Demonstrating the use of if, else and else if ***********
func ifSample() {
	isTheDayBeauty := false
	if isTheDayBeauty {
		fmt.Println("Let's to the beach!")
	} else {
		fmt.Println("Let's sleep!")
	}

	sizeOfYourSpace := "small"
	if sizeOfYourSpace == "small" {
		fmt.Println("The sofa does not fit in your living room")
	} else if sizeOfYourSpace == "Medium" {
		fmt.Println("The sofa fits pretty tight in your living room")
	} else {
		fmt.Println("The sofa fit well in your living room")
	}
}

// ***********  Switch without switch statement ***********
func switchSample() {
	favoriteSport := "Soccer"
	switch favoriteSport {
	case "Basketball":
		fmt.Println("Basketball is in USA")
	case "Volleyball":
		fmt.Println("Volleyball is in Italy")
	case "Soccer":
		fmt.Println("Soccer is in Brazil")
	}
}

// ***********  Logical Operators ***********
func logicSample() {
	//fmt.Println(true && true)
	//fmt.Println(true || true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// ***********  print numbers - 1 to 100 ***********

// islast := false
// for i := 1; i <= 100; i++ {
// 	if i == 100 {
// 		islast = true
// 	}
// 	if !islast {
// 		fmt.Print(i, ", ")
// 	} else {
// 		fmt.Print(i)
// 	}
// }

// ***********  All capital letters unicode code ***********

// for i := 65; i <= 90; i++ { //A-Z (A = 65, B = 66,...)
// 	fmt.Printf("\t%#U", i)
// 	fmt.Println(" - ", i)
// }

// ***********  My birthday year until now and my age ***********

// myYear := 1978 // enter the year of your birthday here
// year := time.Now().Year()
// countAge := -1
// for i := myYear; i <= year; i++ {
// 	countAge++
// 	fmt.Print(i, ", ")
// }
// fmt.Println("My age is:", countAge)
// fmt.Println("My age is:", year-myYear)

// ***********  My birthday year until now and my age 2***********

// myYear := 1978 // enter the year of your birthday here
// countAge := -1
// for myYear <= time.Now().Year() {
// 	countAge++
// 	fmt.Print(myYear, "\t")
// 	myYear++
// }
// fmt.Println("\nMy age is:", countAge)

// ***********  My birthday year until now and my age 3 ***********

// myYear := 1978 // enter the year of your birthday here
// countAge := -1
// for {
// 	fmt.Print(myYear, "\t")
// 	countAge++
// 	myYear++
// 	if myYear > time.Now().Year() {
// 		break
// 	}
// }
// fmt.Println("\nMy age is:", countAge)

// ***********  Remainder after division (modulo function) of numbers between 10 and 100 ***********
// for i := 10; i <= 100; i++ {
// 	fmt.Print(math.Mod(float64(i), 4), "\t")
// }

// ***********  Remainder after division (modulo function) of numbers between 10 and 100 - 2 ***********
// for i := 10; i <= 100; i++ {
// 	fmt.Print((i % 4), "\t")
// }
