package main

import "fmt"

// First letter should be Capital otherwise that property would not be exportable
type User struct {
	Name string
	Email string
	Age int
	Active bool
}

func main()  {
	 shiv := User{"Shiv","shiv@dev.com" ,24, true}
	 fmt.Println(shiv)
	 shiv.GetUser()
	 shiv.NewEmail()
	 fmt.Println(shiv)
}


// We want this method exportable then use First 
// It depends up to you that you want to pass entrire struct or just few properties
// I just called u like an arg for user with Type User
func (u User) GetUser(){
   fmt.Println("What is  user active status:", u.Active)
}

// Here we have manipulate the email it won't change a the actual email property of User because struct is pass as copy of object
// Thats is why pointers come into play
func (u User) NewEmail(){
    u.Email = "test@dev.com"
	fmt.Println("User's new email is:", u.Email)
}


// Here every object is passed as copy of orignal