package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)
func main() {
  myurl := "http://localhost:3000/formdata"

  // Values is a type(struct/interface)
  data := url.Values{}
  data.Add("name", "John Doe")
  data.Add("age", "30")
  data.Add("city", "New York")

  response, err := http.PostForm(myurl, data)

  if err != nil{
	panic(err)
  }

  defer response.Body.Close()

  content, _ := io.ReadAll(response.Body)

  fmt.Println("response content", string(content))
  fmt.Println("response ", response.Header)


}