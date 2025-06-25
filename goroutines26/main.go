package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

//// So we create a goroutine/thread using go keyword,
//// Result will surprise you, it will print only Welcome to Go programming!
//// greeter("Hello, World!") blocks and runs to completion in main goroutine.
//// go greeter("Welcome to Go programming!") starts a goroutine, but main may exit before it runs.
// why? beacuse
// func main() {
// 	go greeter("Hello, World!")  // scheduled concurrently
// 	greeter("Welcome to Go programming!")  // run immediately in main goroutine
// }

// func greeter(s string){
// 	for i:=0; i<5; i++{
// 		fmt.Println(s)
//    }
// }

// now you will see both print as one task is at rest other task has time to run in above code main goroutine runs so fast and stop
// func main() {
// 	go greeter("Hello, World!")
// 	greeter("Welcome to Go programming!")
// }

// func greeter(s string){
// 	for i:=0; i<5; i++{
// 		time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds
// 		fmt.Println(s)
//    }
// }

//// you will all prints running one by one taking time
// func main(){
// 	websites := []string{
//       "http://localhost:4000/api/movies",
// 	  "https://google.com",
// 	  "https://fb.com",
// 	  "https://twitter.com",
// 	}
//     for _, web := range websites{
// 		getStatusCode(web)
// 	}
// }

// // %s -> string, %d-> integer
// func getStatusCode(endpoint string){
// 	res, err := http.Get(endpoint)
//     if err != nil{
// 		panic(err)
// 	}
// 	fmt.Printf("Status code for %s is %d\n", endpoint, res.StatusCode)
// }

//// use go routine with wait groups
//// In Go, we often run multiple goroutines (lightweight threads) in parallel. But we need a way to:
//// ✅ Wait until all those goroutines finish before continuing.

// var wg sync.WaitGroup

// func main(){
// 	websites := []string{
//       "http://localhost:4000/api/movies",
// 	  "https://google.com",
// 	  "https://fb.com",
// 	  "https://twitter.com",
// 	}
//     for _, web := range websites{
// 		wg.Add(1)
// 		// why only 1 but websites are 4?
// 		// because we are adding 1 to the wait group for each goroutine we are starting
// 		go getStatusCode(web)
// 	}
// 	// we genrally put wait at the end here it will tell the main goroutine to wait for all other goroutines to finish
// 	wg.Wait() // Wait for all goroutines to finish
// }

// // %s -> string, %d-> integer
// func getStatusCode(endpoint string){
// 	res, err := http.Get(endpoint)
//     if err != nil{
// 		panic(err)
// 	}
// 	fmt.Printf("Status code for %s is %d\n", endpoint, res.StatusCode)
// 	defer wg.Done() // so this the place where gro routine will tell the wait group that it is done with its work
// }



////All goroutines are calling append(signal, endpoint) at the same time. signal is a shared slice (shared memory).
// Since append() is not atomic, they clash internally (data races). This can cause: Some appends to be lost Corruption of signal Or in your case: only the initial "test" remains visible
// atomic operationa are the operations which happen at the same time in a single step but as we know append involve mupltiple steps like allocating memory, copying data, etc. so it is not atomic
// These steps are not one quick action — so if goroutine A is in step 3 and goroutine B enters step 2, they can step on each other's memory.
// var wg sync.WaitGroup

// var signal  = []string{"test"}

// func main(){
// 	websites := []string{
//       "http://localhost:4000/api/movies",
// 	  "https://google.com",
// 	  "https://fb.com",
// 	  "https://twitter.com",
// 	}
//     for _, web := range websites{
// 		wg.Add(1)
// 		// why only 1 but websites are 4?
// 		// because we are adding 1 to the wait group for each goroutine we are starting
// 		go getStatusCode(web)
// 	}
// 	fmt.Println(signal)
// 	// we genrally put wait at the end here it will tell the main goroutine to wait for all other goroutines to finish
// 	wg.Wait() // Wait for all goroutines to finish
// }

// // %s -> string, %d-> integer
// func getStatusCode(endpoint string){
// 	res, err := http.Get(endpoint)
//     if err != nil{
// 		panic(err)
// 	}
// 	signal= append(signal, endpoint)
// 	fmt.Printf("Status code for %s is %d\n", endpoint, res.StatusCode)
// 	defer wg.Done() // so this the place where gro routine will tell the wait group that it is done with its work
// }


//// solution

var wg sync.WaitGroup
var mut sync.Mutex // Mutex is a mutual exclusion lock that allows only one goroutine to access the shared resource at a time
var signal  = []string{"test"}

func main(){
	websites := []string{
      "http://localhost:4000/api/movies",
	  "https://google.com",
	  "https://fb.com",
	  "https://twitter.com",
	}
    for _, web := range websites{
		wg.Add(1)
		// why only 1 but websites are 4?
		// because we are adding 1 to the wait group for each goroutine we are starting
		go getStatusCode(web)
	}
	// we genrally put wait at the end here it will tell the main goroutine to wait for all other goroutines to finish
	wg.Wait() // Wait for all goroutines to finish, i.e wait all go routines of wait group returns back
	fmt.Println(signal)
}

// %s -> string, %d-> integer
func getStatusCode(endpoint string){
	res, err := http.Get(endpoint)
    if err != nil{
		panic(err)
	}
	mut.Lock()
	signal= append(signal, endpoint)
	mut.Unlock() // unlock the mutex so that other goroutines can access the shared resource
	fmt.Printf("Status code for %s is %d\n", endpoint, res.StatusCode)
	defer wg.Done() // so this the place where gro routine will tell the wait group that it is done with its work
}
