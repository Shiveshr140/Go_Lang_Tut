package main

import (
	"fmt"

	// "golang.org/x/exp/constraints"
)

// func AddInt(a int, b int) int{
// 	sum := a + b
// 	return sum
// }

// func AddFloat(a float64, b float64) float64{
// 	sum := a + b
// 	return sum
// }

// // What if I  pass 3.4, 5.6 i will complain so I need to make another func AddFloat
// func main(){
// 	// var result = AddInt(5, 3)
// 	var result = AddFloat(3.48, 3.6)

// 	fmt.Println("result is:", result)
// }

//// Here comecs generic into play it make code reusable

// func Add[T int | float64]( a T, b T) T{
//    sum := a+b
//    return sum
// }

// func main(){
// 	var result1 = Add(5, 3)
// 	var result2 = Add(3.48, 3.6)

// 	fmt.Printf("result1 is:%v, result2 : %v \n", result1, result2)
// }

//// What if we ewant to use more type it will become little bit verbose so use interface

// func Add[T int | int16 | int8 | float32 | float64]( a T, b T) T{
//    sum := a+b
//    return sum
// }

// type retType interface{
// 	int | int16 | int8 | float32 | float64
// }

// func Add[T retType]( a T, b T) T{
//    sum := a+b
//    return sum
// }

// func main(){
// 	var result1 = Add(5, 3)
// 	var result2 = Add(3.48, 3.6)

// 	fmt.Printf("result1 is:%v, result2 : %v \n", result1, result2)
// }



//// If we do not want to write this intreface we can use order.Constraints => contain union of all number type including string

// func Add[T constraints.Ordered]( a T, b T) T{
//    sum := a+b	
//    return sum
// }

// func main(){
// 	var result1 = Add(5, 3)
// 	var result2 = Add(3.48, 3.6)

// 	fmt.Printf("result1 is:%v, result2 : %v \n", result1, result2)
// }



//// Just for knowledge
//// Basicaly this UserId is alias of type int so if you do not use like T UserId then it will create isue Because that accepts only exactly int or float64 — not types derived from them.  so to solev you need to use ~
//// ~ this will make use to use any alias that has int type
// type UserId int

// func Add[T ~int | float64]( a T, b T) T{
//    sum := a+b	
//    return sum
// }

// func main(){
// 	a := UserId(1)
// 	b := UserId(2)
// 	var result1 = Add(a,b)
// 	var result2 = Add(3.48, 3.6)

// 	fmt.Printf("result1 is:%v, result2 : %v \n", result1, result2)
// }


//// Most common use case

// func mapValues(values []int, mapFunc func(num int) int) []int {
// 	var newValues  []int
// 	for _, value := range values{
// 		newValue := mapFunc(value)
// 		newValues = append(newValues, newValue) 
// 	}

// 	return newValues
// }


// func main(){
// 	result := mapValues([]int{1,2,3}, func(num int) int{
//        return num*2
// 	} )

// 	fmt.Println("result:", result)
// }


//// make it generic

// func mapValues [T constraints.Ordered] (values []T, mapFunc func(num T) T) []T {
// 	var newValues  []T
// 	for _, value := range values{
// 		newValue := mapFunc(value)
// 		newValues = append(newValues, newValue) 
// 	}

// 	return newValues
// }


// func main(){
// 	result := mapValues([]float64{1.002,2.00987,3.28}, func(num float64) float64 {
//        return num*2
// 	} )

// 	fmt.Println("result:", result)
// }



//// Not only with func we can use struct

// type CustomData interface{
// 	constraints.Ordered  | int | float64 | []byte
// }

// type User [T CustomData] struct{
// 	ID int
// 	Name string
// 	Data T
// } 



// func main(){

// 	var user = User[float64]{
//      ID: 4,
// 	 Name: "Shiv",
// 	 Data: 6.80,
// 	}
	
// 	fmt.Printf("result: %+v", user)
// }


//// Use it with map
//// Why Comparable? Because Go requires all map keys to be comparable types — and this constraint ensures only valid types can be used as map keys at compile time.
//// i.e key1(string) == key2(string)
//// int, float64, string, bool       ✅ comparable


type CustomMap [T comparable, V int | string] map[T]V

func main(){
    myMap := make(CustomMap[int, string])

	myMap[3] = "Apple"
	
	fmt.Printf("result: %+v\n", myMap)
}