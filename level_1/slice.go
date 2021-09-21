package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("s - %v, len - %d, cap - %d \n", s, len(s), cap(s))


	// s1  := []int // Error - type []int is not an expression

	// Run time behaviour of slice, it will double its size once
	// its len become euqal to capacity and when you try to add one more element
	// into it
	for  index := range [10]int{0,1,2,3,4,5,6,7,8,9} {
		s = append(s, index)
		fmt.Printf("s - %v, len - %d, cap - %d \n", s, len(s), cap(s))
	}

	fmt.Println("----------------------------------------------")
	// slice copy
	fmt.Println("Slice copy")

	var s2 []int
	copy(s2, s)
	fmt.Printf("s2 - %v, len - %d, cap - %d \n", s2, len(s2), cap(s2))

	fmt.Println("----------------------------------------------")

	s3 := make([]int, len(s)/2)
	fmt.Printf("s3 - %v, len - %d, cap - %d \n", s3, len(s3), cap(s3))
	copy(s3, s)
	fmt.Printf("s3 - %v, len - %d, cap - %d \n", s3, len(s3), cap(s3))
	// create slice with make keyword with desired size 

	fmt.Println("----------------------------------------------")

	s4 := make([]int, len(s))
	fmt.Printf("s4 - %v, len - %d, cap - %d \n", s4, len(s4), cap(s4))
	copy(s4, s)
	fmt.Printf("s4 - %v, len - %d, cap - %d \n", s4, len(s4), cap(s4))

	fmt.Println("----------------------------------------------")

	// s[low_index, high_index) 
	s5 := s[0:len(s)]
	fmt.Printf("s5 - %v, len - %d, cap - %d \n", s5, len(s5), cap(s5))

	s6 := s[0:len(s)/2]
	fmt.Printf("s6 - %v, len - %d, cap - %d \n", s6, len(s6), cap(s6))
}
