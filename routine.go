package main

import (
    "fmt"
    "time"
)

func run(){
	time.Sleep(1 * time.Second)
	fmt.Println("run")
}

func run2(){
	time.Sleep(3 * time.Second)
	fmt.Println("run2")
}
func run3(){
	time.Sleep(5 * time.Second)
	fmt.Println("run3")
}

func add (x int, y int , ch chan int , delay int){
	// channel is used to communicate between goroutines. Here, we are sending the result of x + y to the channel after a delay of 1 second
	time.Sleep(time.Duration(delay) * time.Second)
	
		ch <- x + y 
		// syantx for sending data to channel is channel <- data
}

func main() {
	//cha<- is directional channel, it can only be used to send data to the channel. It cannot be used to receive data from the channel.
	// <-cha is directional channel, it can only be used to receive data from the channel. It cannot be used to send data to the channel.  
	ch := make(chan int)
	ch2 := make(chan int)

	go add(10, 20, ch, 1)
	go add(30, 40, ch2, 2)
	
	for i := 0; i < 2; i++ {	
	select {
	case x := <-ch:
		fmt.Println("Received from channel 1:", x)
	case y := <-ch2:
		fmt.Println("Received from channel 2:", y)	
		
	}
 }
}