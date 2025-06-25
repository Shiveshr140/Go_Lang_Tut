package main

import (
	"fmt"
	"sync"
)

// wg *sync.WaitGroup is useful You're passing wg to functions (especially in larger apps)
// We will never be sure which go will come back first
// func main(){
//    wg := &sync.WaitGroup{} // Pointer to struc
//     score := []int{0}
// 	// wg.Add(3)
// 	wg.Add(1)
// 	go func(wg *sync.WaitGroup){
//       fmt.Println("first routine")
// 	  score = append(score,1)
// 	  wg.Done()
// 	}(wg) 
//     wg.Add(1)
// 	go func(wg *sync.WaitGroup){
// 		fmt.Println("second routine")
// 		score = append(score,2)
//         wg.Done()
// 	}(wg) 
//     wg.Add(1)
// 	go func(wg *sync.WaitGroup){
//       fmt.Println("third routine")
// 	  score = append(score,3)
// 	  wg.Done()
// 	}(wg) 

// 	wg.Wait() // Wait for all goroutines to finish
// 	fmt.Println("Final score:", score)
// }

// in above code everything looks very fine but if run: go run --race you will issues with exit code 66
// This is Where mutex comes into play

// func main(){
//    wg := &sync.WaitGroup{} // Pointer to struc
//    mut := &sync.Mutex{}
//     score := []int{0}
	
// 	wg.Add(3)
	
// 	go func(wg *sync.WaitGroup, m *sync.Mutex){
//       fmt.Println("first routine")
// 	  mut.Lock()
// 	  score = append(score,1)
// 	  mut.Unlock()
// 	  wg.Done()
// 	}(wg, mut) 
   
// 	go func(wg *sync.WaitGroup, m *sync.Mutex){
// 		fmt.Println("second routine")
// 		mut.Lock()
// 		score = append(score,2)
// 		mut.Unlock()
//         wg.Done()
// 	}(wg, mut) 
    
// 	go func(wg *sync.WaitGroup, m *sync.Mutex){
//       fmt.Println("third routine")
//       mut.Lock()	
// 	  score = append(score,3)
// 	  mut.Unlock()
// 	  wg.Done()
// 	}(wg, mut) 

// 	wg.Wait() // Wait for all goroutines to finish
// 	fmt.Println("Final score:", score)
// }


//// Read write mutex, there is rule only apply this where we are reading resource
/// Reading operation is much faster than writing, lock is very necessary when writing but may or may not be necessary when reading
// Read write mutex is used when you have multiple readers and writers so we want that once reading is complete than only allow threads to write
func main(){
   wg := &sync.WaitGroup{} // Pointer to struc
   mut := &sync.RWMutex{}
   score := []int{0}
	
	wg.Add(3)
	
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
      fmt.Println("first routine")
	  mut.Lock()
	  score = append(score,1)
	  mut.Unlock()
	  wg.Done()
	}(wg, mut) 
   
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("second routine")
		mut.Lock()
		score = append(score,2)
		mut.Unlock()
        wg.Done()
	}(wg, mut) 
    
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
      fmt.Println("third routine")
      mut.Lock()	
	  score = append(score,3)
	  mut.Unlock()
	  wg.Done()
	}(wg, mut) 

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
      fmt.Println("third routine")
	  mut.RLock()
      fmt.Println(score)
	  mut.RUnlock()
	  wg.Done()
	}(wg, mut) 

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final score:", score)
}