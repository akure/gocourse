package main

import "fmt"

// collection of field to represent an object that can have a behavior 
type onlineCourse struct {
	name  string
	subject string
	author string
	currency string
	price float32
}

func (oc * onlineCourse)  description() {
	fmt.Println("oc - name = ", oc.name)
	fmt.Println("oc - subject= ", oc.subject)
	fmt.Println("oc - author= ", oc.author)
	fmt.Println("oc - currency = ", oc.currency)
	fmt.Println("oc - price = ", oc.price)

}

// Pointer receiver is normally used for modifying 
// the object 
func (oc * onlineCourse)  setname(n string) {
	oc.name = n
}

func (oc * onlineCourse ) setsubject(s string) {
	oc.subject = s
}

func (oc * onlineCourse ) setauthor(a string) {
	oc.author = a
}

func (oc * onlineCourse ) setcurrency(c  string) {
	oc.currency = c
}

func (oc * onlineCourse)  setprice(p float32) {
	oc.price = p
}

func (oc onlineCourse) description2() {
	// Non Pointer receiver is normally used for
	// using the fields for non-modifying operations
	fmt.Println("oc - name = ", oc.name)
	fmt.Println("oc - subject= ", oc.subject)
	fmt.Println("oc - author= ", oc.author)
	fmt.Println("oc - currency = ", oc.currency)
	fmt.Println("oc - price = ", oc.price)
}
func main() {
	var golevel1 * onlineCourse
	golevel1 = &onlineCourse{}
	//golevel1 := &onlineCourse{}
	golevel1.setname("Go Level #1")
	golevel1.setsubject("Go Programming language")
	golevel1.setauthor("AK")
	golevel1.setcurrency("dollar")
	golevel1.setprice(199.00)
	golevel1.description()
	golevel1.description2()
}
