package main

import "fmt"


// Single ret value
func concat(s1 string, s2 string)  ( r string) {
	return s1 + s2
}

// Multi-Return Value
func  addMulModDiv(x1 int, x2 int) ( add int,  mul int,  div float32, mod int) {
	add = x1 +x2
	mul = x1*x2
	div = float32 (x1) / float32 (x2)
	mod = x1%x2
	return add, mul, div, mod
	// return x1+x2, x1*x2, (float32) x1/ (float32) x2,  x1%x2
}

// Variadic functions - Can pass any number of argument of a
// particular type. ... is an iterable range of elements

func multiconcat( strings ...string)  ( r string) {
	var out string
	for _,  str := range strings {
		out += str
	}
	return out
}
/* Error - can not use ... in return
func  series(num int) ( nums ...int)  {
	return num,num+1, num+2, num+3, num+4
}
*/

func main() {
	out  := concat("hello", ".go")
	fmt.Println("out - %v", out)

	fmt.Println( addMulModDiv(2,3))

	add, mul, div, mod := addMulModDiv(2,3)
	fmt.Printf("add - %d, mul - %d, div - %f, mod - %d \n", add, mul, div, mod)

	fmt.Printf("concat out - %v \n", multiconcat("this", "is", "go"))

	// fmt.Println(series(1))
}
