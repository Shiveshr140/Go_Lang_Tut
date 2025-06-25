package main

import (
	"fmt"
	"io"
	"net/http"
)

var url string = "https://google.com"

func main()  {
	fmt.Println("Google web request!")

	response, err := http.Get(url)

	if err != nil {
      panic(err)
	}

	fmt.Printf("response is of type: %T \n", response)

	defer response.Body.Close()   // it is caller responsibility to close connect after request, better use defer

	dataByte, err := io.ReadAll(response.Body) 

	if err != nil{
		panic(err)
	}

	fmt.Println(string(dataByte))
}

//// you can see the type: *http.Response, which is a pointer remember * so you are not getting the copy of response but u are getting actual response