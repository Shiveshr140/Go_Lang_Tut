package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//// Now if run you will see this err even you have done every thing right
//// strconv.ParseFloat: parsing "3\n": invalid syntax
// Reson is => When you use reader.ReadString('\n'), it does read until the newline (\n). However, the string returned by ReadString('\n') includes the newline character at the end of the input. This is the key point that leads to the issue when trying to convert the input into a number with strconv.ParseFloat().
// func main()  {
// 	fmt.Println("Welcome to my pizza app")
// 	fmt.Println("Please rate out pizza b/w 1 and 5: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input,_ := reader.ReadString('\n')
// 	fmt.Println("Thanks for rating, ", input)

// 	//// now lets add one to the rating, strconv is package
// 	var newRating,err = strconv.ParseFloat(input,64)
// 	if err != nil{
// 		fmt.Println(err)
// 		// panic(err), this will stop the program
// 	}else{
// 		println("now add 1, ", newRating+1)
// 	}
// }


////
func main()  {
	fmt.Println("Welcome to my pizza app")
	fmt.Println("Please rate out pizza b/w 1 and 5: ")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	//// now lets add one to the rating, strconv is package
	var newRating,err = strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		fmt.Println(err)
		// panic(err), this will stop the program
	}else{
		fmt.Println("now add 1: ", newRating+1)
		fmt.Printf("now add 1: %.2f \n: ", newRating+1)
	}
}


    