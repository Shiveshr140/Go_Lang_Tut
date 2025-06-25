package main

import "fmt"

func main() {
 atul := User{"Atul", 19, true}
 fmt.Println(atul)
//  + in %+v print key also
 fmt.Printf("The user detail: %+v\n", atul)
 fmt.Printf("The user name %v and age %v\n", atul.Name, atul.Age)
}

type User struct {
	Name string
	Age  int
	Active bool
}
