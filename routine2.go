package main

import (
	"fmt"
	"time"
	"sync"
)

type Counter struct {
	value int
	lock sync.Mutex // mutex to protect the counter value from concurrent access
}

func count(counter *Counter){
	counter.lock.Lock() // acquire the lock before modifying the counter value
	defer counter.lock.Unlock() // release the lock after modifying the counter value
	counter.value++
	fmt.Println(counter.value)
}

func main() {
	ch := make(chan bool,2) // buffered channel with capacity of 2
	ch <- true // send data to channel
	ch <- true
	<-ch // receive data from channel
	ch <- true // send data to channel
	
	counter := Counter{}

	for i := 0; i < 100; i++ {
		go count(&counter) // pass the address of counter to the goroutine
	}
	time.Sleep(2*time.Second)
}