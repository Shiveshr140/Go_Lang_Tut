package main

import "fmt"

func update(arr1 [3]int){
	arr1[0] = 10
}

func updateAge(arr2 *[3]int){
  arr2[0] = 10 
}

func main()  {
	var fruit [4]string 
    fruit[0] = "apple"
    fruit[1] = "tomato"
    fruit[3] = "peach"
	fmt.Println("See the differnce go by default put space to indicate that something is missing", fruit)
	fmt.Println("length of array will be same as you declare", len(fruit))

	var veg = [3]string{"potato", "ladyfingure", "beans"}
	fmt.Println("So we have declare and intialize array", veg)
	fmt.Println(len(veg))

	// w/o pointer will not modify
	var ages = [3]int{1,2,3}
	update(ages)
	fmt.Println("w/o pointer", ages)

	// with pointer
	updateAge(&ages)
    fmt.Println("with pointer", ages)
	
}