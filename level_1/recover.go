package main

import (
	"fmt"
)

func recoverHelper() {
	println("recoverHelper called.")
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func  badslice() {
	defer recoverHelper()
	s := []int{1, 2, 3}
	fmt.Println(s[100])
 
}

func main() {
	println("start main...")
	badslice()
	println("end main...")

}

