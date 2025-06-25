package Variables2

import "fmt"

var email string = "email.com"

// %T prints the type of the variable
// : wallrus operator
func PrintInfo() {
	var username string = "Shivesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	// no vars style, you can declare the variable w/o using var for that u need wall
	// The fmt.Println() function automatically adds a space between its arguments.
	numOfUsers := 3000
	fmt.Printf("Number of users:%d \n", numOfUsers)
	numOfUsers2 := 3000
	fmt.Println("Number of users:", numOfUsers2)

	fmt.Println("This is the email of the user:", email)

}
