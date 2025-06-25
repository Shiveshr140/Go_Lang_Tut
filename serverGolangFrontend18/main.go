package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {

	response, error := http.Get("http://localhost:3000/")

	if error != nil{
		panic(error)
	}
	fmt.Println("ContentLength is: ", response.ContentLength)
	fmt.Println("Status Code is: ", response.StatusCode)

	content, err := io.ReadAll(response.Body)

	fmt.Println("Content", string(content))
	fmt.Println("Error", err)
}