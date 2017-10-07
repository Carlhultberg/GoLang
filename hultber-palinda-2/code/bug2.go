package main

import "fmt"
import "sync"

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {

	wg := new(sync.WaitGroup)
	wg.Add(11)
	ch := make(chan int)
	go Print(ch,wg)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	wg.Wait()
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg*sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
		wg.Done()
	}
}