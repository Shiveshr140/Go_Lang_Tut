package main

import (
	"fmt"
	// "io"
	"log"
	"net/http"
	"sync"
	"time"
	_ "net/http/pprof" // side-effect import
)

// func fetchUrl(url string, wg *sync.WaitGroup, ch chan string, rateLimiter <- chan time.Time) {
// 	//// it should be at the top because if any return happen before it it will create issue
// 	defer wg.Done()
// 	<- rateLimiter  // block until allowed to proceed
	
// 	res, err1 := http.Get(url)
	
// 	if err1 != nil {
// 		log.Fatal(err1)
// 	}

// 	_, err2 := io.ReadAll(res.Body)

// 	if err2 != nil {
// 		log.Fatal(err2)
// 		ch <- fmt.Sprintf("Error occured while fetching url:", err2)
// 	}

// 	ch <-  fmt.Sprintf("Url: %s, Status: %s \n", url, res.Status)
// 	// fmt.Printf("Content: %s \n ", string(rawData))

// 	defer res.Body.Close()


// }

// func main() {

// 	var wg = &sync.WaitGroup{}
// 	urls := []string{
// 		"https://golang.org",
// 		"https://gobyexample.com",
// 		"https://pkg.go.dev",
// 		"https://httpbin.org/get",
// 	}
// 	var ch = make(chan string, len(urls))  // buffer size = number of goroutines

// 	rate := 20 * time.Second // 1 req per 2 sec
// 	rateLimiter := time.Tick(rate)

// 	for _, url := range urls {
// 		wg.Add(1)
// 		go fetchUrl(url, wg, ch, rateLimiter)
// 	}
	
// 	for msg := range ch{
// 		fmt.Println(msg)
// 	}

// 	close(ch)
  
// 	wg.Wait()

// }


// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10

func ReusableFetch(url string) (res *http.Response, err error){
   res, err = http.Get(url)
	
   
   if err != nil {
	   return res, err
	}
	
	defer res.Body.Close()
	return res, nil

}


func fetchUrl(url string, wg *sync.WaitGroup, ch chan string, rateLimiter <- chan time.Time) {
	//// it should be at the top because if any return happen before it it will create issue
	defer wg.Done()
	<- rateLimiter  // block until allowed to proceed
	
	res,err := ReusableFetch(url)

	if err != nil {
		log.Fatal(err)
		ch <- fmt.Sprintf("Error occurred while fetching url: %v", err)
	}

	ch <-  fmt.Sprintf("Url: %s, Status: %s \n", url, res.Status)
	// fmt.Printf("Content: %s \n ", string(rawData))



}

func main() {
	go func() {
		fmt.Println("pprof server at http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()
	var wg = &sync.WaitGroup{}
	urls := []string{
		"https://golang.org",
		"https://gobyexample.com",
		"https://pkg.go.dev",
		"https://httpbin.org/get",
	}
	var ch = make(chan string, len(urls))  // buffer size = number of goroutines

	rate := 20 * time.Second // 1 req per 2 sec
	rateLimiter := time.Tick(rate)

	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, wg, ch, rateLimiter)
	}
	
	for msg := range ch{
		fmt.Println(msg)
	}

	close(ch)
  
	wg.Wait()

}