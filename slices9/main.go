package main

import (
	"fmt"
	"sort"
)

func main()  {
	// you need to initialize the slices with this syntax
	var fruits = []string{"lemon", "orange"}
	fmt.Printf("The type is: %T \n", fruits)

	// as slices are dynamic we can add elements
	fruits = append(fruits, "mango","apple")
	fmt.Println(fruits)

	fruits1 := fruits[1:]
	fmt.Println(fruits1)

	// make built in method, Use it when You need to create a slice, map, or channel with a specific initial size or capacity.
	var highScores = make([]int, 4)
	highScores[0] = 233
	highScores[1] = 133
	highScores[2] = 333
	highScores[3] = 433
	// highScores[4] = 233 will show an error like index out of range
    fmt.Println(highScores)

    highScores = append(highScores, 789, 1010 )
	fmt.Println("You can make it dynamic using append", highScores)

	// sorting
	alphabets := make([]string, 4)
	alphabets[0] = "A"
	alphabets[1] = "G"
	alphabets[2] = "I"
	alphabets[3] = "D"
	sort.Strings(alphabets)
	fmt.Println(alphabets)

	//// remove element
	var courses = []string{"javascript", "react", "swift", "node", "python"} 
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...) 
	fmt.Println(courses)
}