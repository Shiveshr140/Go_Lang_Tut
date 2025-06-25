package main

import "fmt"

// When a function executes it executes line by line even though it is compiled because that is how flow works,
// As soon as you mark defer keyword whatever the next line you marked as execution that is going to execute at very end of the function
// func main()  {
	// 	defer fmt.Println("World")
	// 	fmt.Println("Hello World")
	// }
	
	
// multiple defer execute last in first out
func main()  {
	defer fmt.Println("last")
	defer fmt.Println("second")
	defer fmt.Println("first")
	fmt.Println("Hello World")
	myDefer()
}



func myDefer(){
	for i := 0; i < 5; i++ {
       defer fmt.Println(i)
	}
}


// Hello
// myDefer() => it will create a stack as defer is used [0,1,2,3,4] => follows LIFO
// first 
// second 
// last