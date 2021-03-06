package main

import "time"
import "fmt"

func Remind(text string, delay time.Duration) {

	
	for {
		t := time.Tick(delay)
		time.Sleep(delay)
		select {

		case now := <-t:
			fmt.Print("\r Klockan är ", now.Hour(), ":", now.Minute(),":", now.Second(), " ", text)
		}
	}
}

func main() {

	go Remind("Dags att äta", 3 * time.Second)
	go Remind("Dags att arbeta", 8 * time.Second)
	go Remind("Dags att sova", 24 * time.Second)
 	select {}
}