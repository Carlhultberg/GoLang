package main

import (
	"fmt"
	"sync"
)
const NUM = 2
// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
	sum := 0
	for _, value := range a {
		sum += value
	}
	res <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)
	wg := new(sync.WaitGroup)
	count := 0
	total := 0
	wg.Add(NUM)
	go func() {
		for val := range ch {
			count =+1
			fmt.Printf("%d", val)
			total += val
			if count != NUM {
				fmt.Print(" + ")
			}
			wg.Done()
		}
	}()
	wg.Wait()
	close(ch)
	fmt.Printf(" = %d", total)



	// TODO: Get the subtotals from the channel and print their sum.
}
