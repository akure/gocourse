package main

import "fmt"

func main() {

	// Example of if
	if true {
		fmt.Println("true")
	}

	// Example of if - else 
	var a bool = true
	var b bool = false

	if  a == b {

		fmt.Println("a==b")

	} else  {

		fmt.Println("a!=b")

	}

	// Example of if -else if else 
	var c1 int = 10
	var c2 int = 11
	var c3 int = 12

	if c1 == 12 {
		fmt.Println("c1 == 12 branch")
	} else if c2 == 12 {
		fmt.Println("c2 == 12 branch")
	} else if c3 == 12 {
		fmt.Println("c3 == 12 branch")
	} else  {

		fmt.Println("c1, c2 and c3, don't match with 12")
	}

}
