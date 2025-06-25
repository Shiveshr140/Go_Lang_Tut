package main

import "fmt"

type Company struct{
	Name string
	Employes int
	Revenue int
}

func main(){
	kodnest := Company{Name:"Kodnest", Employes: 70, Revenue:70000000}
	updateStruct(&kodnest)
	fmt.Printf("kodnest struct has type: %T, kodnest struct value is %+v\n ",kodnest, kodnest)
}

func updateStruct(c *Company){
  c.Name= "Goggle"
}

// func updateArrWithPointer(arr *[]string){
// 	*arr = append(*arr,"d")
// }