package main

import "fmt"

func main() {

		var x1 *int
		var x2 = new(int)

		//  fmt.Println("x1 - ", x1, ", *x1 - ", *x1)
		fmt.Println("x1 - ", x1) // nil , *x1 -> panic
		fmt.Println("x2 - ", x2, ", *x2 - ", *x2)

		s := new([]string)
		fmt.Println("s - " , s == nil) // false
		fmt.Println("s len -" ,  len(*s))  // 0
		fmt.Println(*s == nil) // true

	    m := new(map[string]int)
	    fmt.Println("m  - ",m == nil) // false
		fmt.Println(*m == nil) // true

		c := new(chan int)
		fmt.Println("c - ", c == nil) // false
		fmt.Println(*c == nil) // true

}
