package main

import (
 "fmt"
 "unsafe"
)

type empty interface {
}


func main() {
	var v interface{}

	var v1 empty

	fmt.Printf("V  --- v = %v , T = %T, size = %d \n",v,v, unsafe.Sizeof(v))
	fmt.Printf("V1 --- v = %v , T = %T, size = %d \n",v1,v1, unsafe.Sizeof(v1))
}
