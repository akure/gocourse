package main

import "fmt"

type EmbeddedInterfaceType interface {
	test()  bool

}

type EmbeddedType1 struct {
	name string
}

func (e EmbeddedType1)  test() bool {
	// Printing the type and object on which it is called
	fmt.Printf("test() ===>  e - Type = %T, Value =  %v \n", e, e)
	return true
}

type EmbeddedType2 struct {
	name string
}

func (e EmbeddedType2)  test() bool {
	// Printing the type and object on which it is called
	fmt.Printf("test() ===>  e - Type = %T, Value =  %v \n", e, e)
	fmt.Printf("test() ===>  e - Type = %T, Value =  %v \n ", e, e)
	return true
}

type EmbeddingType struct {
	ei EmbeddedInterfaceType
}


func main() {
	var ei1 EmbeddedInterfaceType
	var e1 EmbeddedType1
	var e2 EmbeddedType2
	var o  EmbeddingType

	fmt.Printf("ei1 - Type = %T, Value =  %v \n", ei1, ei1)
	fmt.Printf("e1 - Type = %T, Value =  %v \n", e1, e1)
	fmt.Printf("e2 - Type = %T, Value =  %v \n", e2, e2)
	fmt.Printf("o - Type = %T, Value =  %v \n", o, o)

	// ei2 :=  EmbeddedInterfaceType{} // invalid composite literal type EmbeddedInterfaceType
	// fmt.Printf("ei2 - Type = %T, Value =  %v \n", ei2, ei2)

	var ei2 EmbeddedInterfaceType
	fmt.Printf("ei2 - Type = %T, Value =  %v \n", ei2, ei2)

	e1.name = "rocky"
	ei2 = e1 // change the type of ei2 from nil to the type of e1. 
	// ei2 becomes a variable of type EmbeddedType1
	// ei2 is representing e1
	fmt.Printf("ei2 - Type = %T, Value =  %v \n", ei2, ei2)
	ei2.test()


	e2.name = "Matrix"
	ei2 = e2  // change the type of ei2 from EmbeddedType1 to EmbeddedType2
	// ei2 becomes a variable of type EmbeddedType2
	// ei2 is representing e2
	fmt.Printf("ei2 - Type = %T, Value =  %v \n", ei2, ei2)
	ei2.test()
	// Methods 
}
