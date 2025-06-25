package main

import (
	"math/rand"
	"fmt"
)

func main() {
	var loginCount int = 11
	var result string
	if loginCount > 10 {
		result = "regular user"
	} else if loginCount < 10 {
		result = "rarely visitor"
	} else {
		result = "user visited 10 times"
	}

	fmt.Println(result)

	// initiaze and check at the same time same time syntax
	if rate := 2000; rate < 1000 {
		fmt.Println("Rate is low")
	} else {
		fmt.Println("Rate is high")
	}

	//   Switch case
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The value of dice:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("your count is 1", diceNumber)
	case 2:
		fmt.Println("your count is 2", diceNumber)
	case 3:
		fmt.Println("your count is 3", diceNumber)
		
	default:
		fmt.Println("what is this!")

	}

}
