package main

import ( 
	"fmt"
	"time"
 )

func printData(c chan *int) {
   // for {
   time.Sleep(time.Second * 3)
	
   fmt.Println("------go routine-------")
   data := <-c
   fmt.Println("Data in channel is: *data ", *data, ", data - ", data)
	// }
   fmt.Println("-----Exiting goroutine -------")
}

func main() {
   fmt.Println("Main started...")
   var a = 10
   b := &a
   //create channel
   c := make(chan *int)
   go printData(c)
   fmt.Println("Value of b before putting into channel", *b, ", &a - ", &a, ", b - ", b)
   c <- b

   fmt.Println("------go main routine-------")
   a = 20
   fmt.Println("Updated value of a: ", a)
   fmt.Println("Updated value of b: ", *b)
   time.Sleep(time.Second * 2)
   fmt.Println("Main ended...")
}


