package main

import "fmt"

func main() {
    var i  int
	fmt.Println("i - ", i)
	i = 10
	fmt.Println("i - ", i)

	var b bool
	fmt.Println("b - ", b)

	b = true
	fmt.Println("b - ", b)

	var f float32

	fmt.Println("f - ", f)
	f = 10.10
	fmt.Println("f - ", f)

	var s string
	fmt.Println("s -", s)

	s = "go string"
	fmt.Println("s -", s) 
}
