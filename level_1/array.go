package main

import "fmt"

func main() {
	const size int = 10
	var arr[size] int

	fmt.Printf("arr - %v, len - %d \n", arr, len(arr))

	// setters
	for  index, _ := range arr {
		arr[index] = index
	}

	fmt.Printf("arr - %v \n", arr)

	const size2 = 5 
	arr2 := [size2]int{0,1,2,3,4}
	fmt.Printf("arr2 - %v, len - %d \n", arr2, len(arr2))

}
