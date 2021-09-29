package main

import "fmt"

func printChanData(c chan int ) {
    // <-c is a way to deque data from the channel. 
	tmp :=  <-c
	fmt.Println("c data - ", tmp )
}


func main() {
		//create unbuffered channel of type int
		c1 := make(chan int)
		// call to goroutine
		go printChanData(c1)
		// put the data in channel
		c1 <- 10
		fmt.Println("Exit...")
}
