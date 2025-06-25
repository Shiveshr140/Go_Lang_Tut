package main

import (
	"fmt"
	"io"

	"os"
)

// Panic will stop the execution and show the error
// by default we will have os package which can be used for creating file
// for writting anything we will have io package
func main(){
	content := "This needs to go in file - myText1.txt"

	file, error := os.Create("./myText1")

	if error != nil {
		panic(error)
	}
    
	var length, err = io.WriteString(file, content)
    
	if err !=nil{
		panic(err)
	}

	fmt.Println("Length is:", length)

	defer file.Close()

	readFile("./myText1")
}

//  creating a file is done by os package but for reading and doing manipulation we need ioutil
// When we read data will never read in string format this thing apply when yoy read the data from internet data comes in byte format
func readFile(fileName string){
//    dataByte, err :=  ioutil.ReadFile(fileName)
   dataByte, err := os.ReadFile(fileName)

   if err != nil{
	panic(err)
   }
  
   fmt.Println("This file content in databyte: \n", dataByte)
   
   fmt.Println("This file content is: \n", string(dataByte))

}