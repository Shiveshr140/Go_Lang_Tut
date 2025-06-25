package main

import (
	"fmt"
	"sync"
)

// As we can use make to create any datatype here we will use to create channel
// as channels are they way to communicate between goroutines we need define what type of data we want to send through the channel
// and we can use make to create a channel of that type
// If we run this we will see famous error "all goroutines are asleep - deadlock!" beacuse channel only allowing to pass a value only if somebody is listening to it.
// Now you can think guy who is listening to the channel is at one line below the line where we are sending a value to the channel.
// If we swap them still we will the same error
// Here is the sold reason  both are running on same goroutine that is main as we need another go routine to listen
// func main(){
//   fmt.Println("Channel and Deadlock Example")
//   myCh := make(chan int) // Create a channel of type int
//   myCh <- 5 // <- it is a way Send/add a value to the channel
//   fmt.Println("Value sent to channel:", <-myCh) // <- it is a way to receive a value from the channel

// }

//// this will also not work because we are trying to send a value to the channel before the goroutine is created to listen to it
// func main(){
//   fmt.Println("Channel and Deadlock Example")
//   myCh := make(chan int)
//   myCh <- 5
//   go func(){
//     fmt.Println("Value sent to channel:", <-myCh)
//   }()

// }

//// This will work because we are creating a goroutine to listen to the channel before sending a value to it

// func main(){
//   fmt.Println("Channel and Deadlock Example")
//   myCh := make(chan int)
//   go func(){
// 	  myCh <- 5
//    }()

//   fmt.Println("Value sent to channel:", <-myCh)
// }

//// Proper way to use channel is to create a goroutine to listen to the channel and then send a value to it

// func main(){
//   wg := &sync.WaitGroup{}
//   fmt.Println("Channel and Deadlock Example")
//   myCh := make(chan int)
//   wg.Add(2)
//   go func(c chan int, wg *sync.WaitGroup){
// 	  fmt.Println("Goroutine is ready to receive value from channel")
// 	  fmt.Println("Value sent to channel:", <-myCh)
// 	  wg.Done()
//    }(myCh, wg)

//    go func(c chan int, wg *sync.WaitGroup){
// 	  myCh <- 5
// 	  wg.Done()
//    }(myCh, wg)

//   wg.Wait()
// }

//// If you add two add into channel you have two add two recievers
//// 2 is buffered channel which means it can hold two values before it blocks
// make(chan int) — unbuffered, only works when both sender and receiver are ready at the same time.
// make(chan int, 2) — buffered, allows the sender to drop 2 items in the pipe even if nobody is listening yet.

// func main(){
//   wg := &sync.WaitGroup{}
//   fmt.Println("Channel and Deadlock Example")
// //   myCh := make(chan int)
//   myCh := make(chan int, 2)
//   wg.Add(2)
//   go func(c chan int, wg *sync.WaitGroup){
// 	  fmt.Println("Goroutine is ready to receive value from channel")
// 	  fmt.Println("Value1 sent to channel:", <-myCh)
// 	//   fmt.Println("Value2 sent to channel:", <-myCh)
// 	  wg.Done()
//    }(myCh, wg)

//    go func(c chan int, wg *sync.WaitGroup){
// 	  myCh <- 5
// 	  myCh <- 6
// 	  wg.Done()
//    }(myCh, wg)

//   wg.Wait()
// }

//// You can close the channel to indicate that no more values will be sent on it. This is useful for signaling completion to the receiver.

// func main(){
//   wg := &sync.WaitGroup{}
//   fmt.Println("Channel and Deadlock Example")
// //   myCh := make(chan int)
//   myCh := make(chan int, 2)
//   wg.Add(2)
//   go func(c chan int, wg *sync.WaitGroup){
// 	  fmt.Println("Goroutine is ready to receive value from channel")
// 	  fmt.Println("Value1 sent to channel:", <-myCh)
// 	//   fmt.Println("Value2 sent to channel:", <-myCh)
// 	  wg.Done()
//    }(myCh, wg)

//    go func(c chan int, wg *sync.WaitGroup){
// 	//   close(myCh) it will show an error as we are sending value to the channel after closing it
// 	  myCh <- 5
// 	  myCh <- 6
// 	  close(myCh)
// 	  wg.Done()
//    }(myCh, wg)

//   wg.Wait()
// }

// //// Listening to close channel is not an issue, it will return 0
// func main(){
//   wg := &sync.WaitGroup{}
//   fmt.Println("Channel and Deadlock Example")
// //   myCh := make(chan int)
//   myCh := make(chan int, 2)
//   wg.Add(2)
//   go func(c chan int, wg *sync.WaitGroup){
// 	  fmt.Println("Goroutine is ready to receive value from channel")
// 	  fmt.Println("Value1 sent to channel:", <-myCh)
// 	//   fmt.Println("Value2 sent to channel:", <-myCh)
// 	  wg.Done()
//    }(myCh, wg)

//    go func(c chan int, wg *sync.WaitGroup){
// 	   close(myCh)

// 	//   myCh <- 5
// 	//   myCh <- 6
// 	  wg.Done()
//    }(myCh, wg)

//   wg.Wait()
// }

// // But this 0 will an issue if we try to send 0 by check <- 0 we won't whether it is comming from by channel close or we are sending 0
// Ok what if you use the close() the cannel even before listening this will create a panic go has solution to mark arrow <- to make goroutine either send ONLY or r ONLY
func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("Channel and Deadlock Example")

	myCh := make(chan int, 2)
	wg.Add(2)

	// r ONLY channel
	go func(c <- chan int, wg *sync.WaitGroup) {
		close(myCh)
		// you can check if channel is closed by checking the second value returned by the receive operation
		val, isChannelOpen := <- myCh
		fmt.Println("Is channel open?", isChannelOpen)
		fmt.Println("Value recieved from channel:", val)
		wg.Done()
	}(myCh, wg)

   // send ONLY 
	go func(c chan<- int, wg *sync.WaitGroup) {
		// myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
