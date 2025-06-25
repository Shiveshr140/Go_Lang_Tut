package main

import "fmt"

/// you may wonder we did not call this main function like main() still go tool or compiler knows about it so this is the proof that 
//// this main function is entry point. 
//// You can not write a function inside the fnnction but can call them
func main()  {
   fmt.Println("Welcome to functions in go-lang")
   greet()
   result := adder(5, 9)
   fmt.Println("result of adding value1 and value2:", result)

   proResult, _ := proAdder(1,2,3,5)
   fmt.Println("total value", proResult)
}

func adder(val1 int, val2 int) int  {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
   total := 0
   for _,value := range values{
	 total += value
   }
   
   return total, "Hi proAdder function"
}

func greet(){
	fmt.Println("hello user!")
}


