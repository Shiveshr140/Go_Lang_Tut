package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	myUrl := "http://localhost:3000/"
	requestBody := strings.NewReader(`
    {
      "name": "John Doe",
	  "age": "30"
	}
 `)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// response.Body is a stream (io.Reader)
	content, _ := io.ReadAll(response.Body)
   
	fmt.Println("response:", string(content))
	
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	fmt.Println("response :", response.Body)

}
