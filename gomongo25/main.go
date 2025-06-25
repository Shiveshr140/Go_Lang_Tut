// aplication which is designed in go is expecting that one go file at root directory and that can be  main.go
// it is recommecded to open it as a workspace in vscode if your using multiple files

package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/Shiveshr140/gomongo/routers"
)


func main() {	
	fmt.Println("Hello MongoDB World!")
	
	log.Fatal(http.ListenAndServe(":4000", router.Router()))
}