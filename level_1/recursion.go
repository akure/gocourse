package main

import "fmt"

func factorial(n int) int {
	fmt.Println("factorial n - ", n)
	if n == 0  { return 0 }
	if n == 1 { return 1 }
	fact := n * factorial(n-1)
	fmt.Println("factorial n - " , n , " => ", fact )
	return fact
	// return n * factorial(n-1)
}

func main() {
	fmt.Println("main - ", factorial(5))

}
