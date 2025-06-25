package main

import "fmt"
 


func main(){
	languages := make(map[string]string)
	languages["py"] = "python"
	languages["js"] = "javascript"
	languages["rb"] = "ruby"

	fmt.Println(languages)
	fmt.Println("js shorts for:", languages["js"])

	// Delete
	delete(languages,"js")
    fmt.Println("after deleting js", languages)

	// for loop with map
	for key,value := range languages{
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}	