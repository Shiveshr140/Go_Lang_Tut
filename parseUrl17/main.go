package main

import "fmt"
import "net/url"


func main()  {
	myUrl := "https://www.google.com/search?q=golang&source=hp"
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme) // protocol
	fmt.Println(result.Host)   // Domain + port
	fmt.Println(result.Port()) // if port in url then it will show 
	fmt.Println(result.RawQuery)

    qparams := result.Query()
    fmt.Printf("type of qparams: %T \n", qparams)   // url.Values => gives the idea of key value pair
	fmt.Println(qparams["q"])

	// for _, val := range qparams{
	// 	fmt.Println("Params is: ", val)
	// }

	// reverse order, remember use reference as pointer 
	// & operator creates a pointer to a new url.URL struct.
    // So, partsOfUrl is a pointer to a url.URL, not a direct value.
	partsOfUrl := &url.URL{
		Scheme:"https",
		Host: "www.gogle.com",
		Path:"/",
		RawPath: "q=golang&source=hp",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
