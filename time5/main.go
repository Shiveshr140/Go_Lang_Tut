package main

import (
	"fmt"
	"time"
)

// Why do we use 01-02-2006 instead of something like YYYY-MM-DD?
// In Go, when you want to format a date/time with Time.Format(), you don’t use the usual date/time format specifiers like YYYY, MM, DD, etc. Instead, Go has a very specific reference date that you need to use to build the format string.

// The Reference Date:
// The reference date in Go is January 2, 2006 at 3:04:05 PM.

// This date is significant because Go uses this specific date as the model to determine how to format and parse time. Here’s how it works:

// Year: 2006

// Month: 01

// Day: 02

// Hour: 03

// Minute: 04

// Second: 05

// AM/PM: PM



func main()  {
	fmt.Println("Welcome to the time study in golang")
	var parseTime = time.Now()
	fmt.Println("Time is: ", parseTime)

	////It will print date only, why 01-02-2006, it is pretty standard in go you have to use this syntax
	fmt.Println("Print date only", parseTime.Format("01-02-2006"))
	fmt.Println("Print date only", parseTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2025, time.August, 3, 10, 20, 20, 10, time.UTC)

	fmt.Println("We have created a date: ", createdDate)

	fmt.Println("Lets make it pretty: ", createdDate.Format("01-02-2006 Monday"))

}