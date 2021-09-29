package main

import (
  "fmt"
)
// buffered go routines are not synchronised by default
func loop(ch chan int) {
  // time.Sleep(time.Second)
  fmt.Println("loop entered\n")

  for i := range ch {
    fmt.Println(i)
    // time.Sleep(time.Second)
  }
}

func main() {
  // only modify this line to defined the capacity
  ch := make(chan int, 3)
  defer close(ch)

  go access(ch)

  for i := 0; i < 100; i++ {
    ch <- i
    fmt.Println(" Element: ", i, " , len - ", len(ch) , ", cap - " , cap(ch))
  }

  // time.Sleep(3 * time.Second)
}
