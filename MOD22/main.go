package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func main(){
	Greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
    
	log.Fatal(http.ListenAndServe(":3005", r))
}

func Greeter(){
	fmt.Println("Hello, Kodnest!")
}

// why ResponseWriter and Request are passed as parameters?
// ResponseWriter is used to construct the HTTP response that will be sent back to the client.
//  ResponseWriter is a interface provides methods to set the HTTP status code, headers, and write the response body.
// Why pointer angainst Request?
// Request is a pointer to the http.Request struct, which contains information about the incoming HTTP request.
// Using a pointer allows the handler to modify the request if needed, such as reading form data or query parameters.
func ServeHome(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte("<h1>Welcome to Kodnest</h1>"))	
}