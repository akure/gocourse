package main

import (
  "fmt"
  "time"
)


// buffered go routines are not synchronised by default
func loop(ch chan int) {

    fmt.Println("start infinit loop in this go routine\n")
    for {
        time.Sleep(time.Second)
	}

}

func main() {
  ch := make(chan int, 3)
  defer close(ch)
  go loop(ch)

 for i := 0; i < 100; i++ {
    ch <- i // This program will block here, after writing three elements into it.
    fmt.Println("Element : ", i, " , len - ", len(ch) , ", cap - " , cap(ch))
  }
}
