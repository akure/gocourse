package main

import  "fmt"

func swap (iptr * int, jptr * int)  {
	tmp := *iptr
	*iptr = *jptr
	*jptr = tmp
}

func main() {
	var i  int

	iptr := &i
	*iptr = 100
	// Value of &i and iptr are same; Address of i
	// Value of i and *iptr are same; Value at address of i
	// Value of &iptr is the address of variable iptr
	fmt.Println("i = ", i, ", &i = ", &i , ", iptr = ", iptr, ", &iptr = ", &iptr, " *iptr = ", *iptr)

	x1 := 10
	x2 := 20
	swap(&x1, &x2)

	fmt.Println("x1 = ", x1, ", x2 = ", x2)

	// swap(x1, x2) // cannot use x1 (type int) as type *int in argument to swap, cannot use x2 (type int) as type *int in argument to swap
}
