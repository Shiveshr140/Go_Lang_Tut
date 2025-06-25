package main

import "fmt"

func update(n int){
		n=99
}

func updateAge(x *int){
	*x = 99
}


func main(){
	fmt.Println("Welcome to pointers")

	var pointer *int 
	fmt.Println("Value of this pointer", pointer)

	number := 23
	var myPointer = &number

	fmt.Println("myPointer is refernce of memory address of number using & which is just a reference", myPointer)
	fmt.Println("Actual value that is stored in that memory address using *", *myPointer)

	// Lets see why it is usefull
	*myPointer = *myPointer + 2
	fmt.Println("Operations are performed on actual value not copies of those values", *myPointer)

	// w/o pointer
	var age int = 30
	update(age)
	fmt.Println("Age won't age as we are passing copy of age", age)
	 
	// with pointer
	myAge := 25
    updateAge(&myAge)
	fmt.Println("Age will be change as we are passing actual age not copy", myAge)


	//// Another e.g
		// declaring a pointer
	var newPointer *string
	newStr := "hi"
	// assiging address to the pointer
	newPointer = &newStr
    // derefrencing pointer to get value
	fmt.Println(*newPointer)

}



