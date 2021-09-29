package main

import "fmt"

func main() {

    m :=make(map[string]string)
    m1 :=make(map[string]string, 100)

	fmt.Println("m map len =", len(m))
	fmt.Println("m1 map len =", len(m1)) // it will have enough internal space

    s :=make([]int, 2)
    s1 :=make([]int, 2,100)
	fmt.Println("s len =", len(s), ", cap = ", cap(s))
	fmt.Println("s1 len =", len(s1), ", cap = ", cap(s1))

	c :=make(chan int)
    c1 :=make(chan int,100)
	fmt.Println("c len =", len(c), ", cap = ", cap(c))
	fmt.Println("c1 len =", len(c1), ", cap = ", cap(c1))
    // c <-100

}
