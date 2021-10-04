package main

import "fmt"

type EmbeddedType1 struct {
	name string
}

type EmbeddedType2 struct {
	name string
}

type EmbeddedType3 struct {
	name string
}

func (e EmbeddedType1) test() {
	fmt.Printf("test() ===> type = %T , value = %v \n", e, e)
}

/*
// This method will produce error - ambiguous selector o2.test
func (e EmbeddedType2) test() {
	fmt.Printf("test() ===> type = %T , value = %v \n", e, e)
}
*/

type EmbeddingType struct {
	EmbeddedType1
	EmbeddedType2
	EmbeddedType3
}

func main() {
	var e1 EmbeddedType1 = EmbeddedType1{"first"}
	var e2 EmbeddedType2 = EmbeddedType2{"Second"}
	var e3 EmbeddedType3 = EmbeddedType3{"Third"}

	var o1 EmbeddingType

	fmt.Println("e1 - ", e1)
	fmt.Println("e2 - ", e2)
	fmt.Println("e3 - ", e3)

	fmt.Println("o1 - ", o1) // o1 -  {{} {} {}}

	o2  := EmbeddingType{ EmbeddedType1{"first"} , EmbeddedType2{"second"}, EmbeddedType3{"third"} }
	fmt.Println("o2 - ", o2) // o1 -  {{first} {second} {third}}

	o2.test()
}

