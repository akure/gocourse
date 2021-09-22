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


func main() {
	var golevel1 onlineCourse
	golevel1.name = "Go Level #1 "
	golevel1.subject = "Go programming language"
	golevel1.author = "AK"
	golevel1.currency = "dollar"
	golevel1.price = 199.00

	fmt.Println("onlineCourse = ", golevel1)


	// object initialization with struct literal zero value
	course1  :=  onlineCourse{}
	fmt.Println("onlineCourse = ", course1)

	// object initialization with struct literal field values
	// in the order of their presence
	course2  :=  onlineCourse{"Go Level #1", "Go Programming language", "AK", "dollar" , 199.00 }
	fmt.Println("onlineCourse = ", course2)

	// object initialization with struct literal field values
	// order of fields does not matter, also you can skip any particular field based on
	// your requirement
	course3  :=  onlineCourse{author: "AK", currency :"dollar" , price :199.00, name : "Go Level #1", subject : "Go Programming language" }
	fmt.Println("onlineCourse = ", course3)

	// object initialization with struct literal field values
	// order of fields does not matter, also you can skip any particular field based on
	// your requirement. Notice the last field "," and }
	// Rule -> last , in the struct literal is optional if , and } are on the same line. 
	course3  :=  onlineCourse{author: "AK", currency :"dollar" , price :199.00, name : "Go Level #1", 
	subject : "Go Programming language",
	}

	fmt.Println("onlineCourse = ", course3)

}
