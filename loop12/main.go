package main

import "fmt"

func main()  {
	// days := []string{"sunday", "monday"}

	// for i:=0; i<len(days); i++{
	// 	fmt.Println("today is:", days[i])
	// }

	//// unlike other language in go here i will not refer to actual value but index

	// for i := range days{
	// 	fmt.Println("Today is:", days[i])
	// }

	//// for actual value
	// for index, day := range days{
	// 	fmt.Printf("Today is %v, the index is %v\n", day, index )
	// }

	//// while loop

	// rogoue := 1
	// for rogoue < 10{
    //   fmt.Println("rogoue is:", rogoue)
	//   rogoue++ 
	// }
	// for rogoue < 10{
	//   if rogoue == 5 {
	// 	break
	//   }
    //   fmt.Println("rogoue is:", rogoue)
	//   rogoue++ 
	// }

	// rogoue := 1

	// for rogoue < 10{
	//   if rogoue == 5 {
	// 	rogoue++
	// 	continue
	//   }
    //   fmt.Println("rogoue is:", rogoue)
	//   rogoue++ 
	// }

	rogoue := 1

	for rogoue < 10{
	  if rogoue == 5 {
		goto lco
	  }
      fmt.Println("rogoue is:", rogoue)
	  rogoue++ 
	}

	lco:
	     fmt.Printf("Jumpying to kodnest.com")
}
