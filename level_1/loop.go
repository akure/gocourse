package main

import "fmt"
import "time"

func main() {

	// Loop #1 
	fmt.Println("Loop #1 - Whilw type of for loop")
	i := 0
    for i <= 10 {
        fmt.Println(i)
        i ++
    }

	// Loop #2
	fmt.Println("Loop #2 - C like three part loop")
	for j := 0 ; j < 10; j++ {
	fmt.Println(j)
	}

	// Loop #3
	fmt.Println("Loop #3 - Do-While type of for loop")
	// It will execute the for block for atleast one time
	for ok := true; ok; ok = false {
      fmt.Println(ok)
	}

	fmt.Println("Loop #4 - range based for loop ")

	// slice of words
	strings := []string{"this", "is", "go"}
	for index, word := range strings {
		fmt.Println(index, word)
	}

	fmt.Println("Loop #5 - range based for loop, only index")
	for index, _ := range strings {
		fmt.Println(index)
	}

	fmt.Println("Loop #6 - range based for loop, only value")
	for _, value := range strings {
		fmt.Println(value)
	}

	
	// Map Iteration
	m := map[string]int{
    "first":  1,
    "second": 2,
    "third": 3,
	}

	fmt.Println("Loop 7 - range based for loop for map, both k v")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Loop 8 - range based for loop for map, only key")
	for k, _ := range m {
		fmt.Println(k)
	}


	fmt.Println("Loop 9 - range based for loop for map, only value")
	for _, v := range m {
		fmt.Println(v)
	}

	// Channel iteration

	ch := make(chan string)
	// launch a goroutine on the fly
	go func() {
    ch <- "this"
	fmt.Println("ch word 1...")
    ch <- "is"
	fmt.Println("ch word 2...")
    ch <- "go"
	fmt.Println("ch word 3...")
    close(ch)
	fmt.Println("ch closed")
	}()

	fmt.Println("Loop 10 - range based for loop for channel")

	for word := range ch {
		time.Sleep(1*time.Second)
		fmt.Println(word)
	}

	// Loop #N
	fmt.Println("Loop # - Infinit Loop")
	for  { }
}
