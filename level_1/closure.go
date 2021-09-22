package main

import "fmt"

// Function that returns another function ( closure ) 
func initAndReturnSeqClosure()  func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

func main() {

	fmt.Println("First -")
	// Create a new sequence
	seq := initAndReturnSeqClosure()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

	fmt.Println("Second -")
	// Create a new sequence
	seq1 := initAndReturnSeqClosure()
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())

	fmt.Println("Third -")
	// Create a new sequence
	seq1 = initAndReturnSeqClosure()
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())

	fmt.Println("Fourth -") 
	// Resume from the previous saved value
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
}
