package main

import "fmt"

func printChanData(c chan int ) {
	for {
		// <-c is a way to deque data from the channel. 
		tmp :=  <-c
		fmt.Println("c data - ", tmp )
	}
}



func main() {
		//create unbuffered channel of type int
		c1 := make(chan int, 10)
		// call to goroutine
		go printChanData(c1)
		// put the data in channel
		for i := 0; i < 100 ; i++ { 
		c1 <- i
		}
		fmt.Println("Exit...")
}
