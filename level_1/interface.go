package main

import (
    "fmt"
    "strconv"
    "log"
)
// Stringer interface is defined in fmt package
//  type Stringer interface {
//	 String() string
// }
type onlineCourse struct {
    name  string
    subject string
    author string
    currency string
    price float32
}


// onlineCourse type satisfies the stringer interface
func (oc onlineCourse) String() string {
    return fmt.Sprintf("onlineCourse: %s - %s - %s - %s - %f", oc.name, oc.subject, oc.author, oc.currency, oc.price)
}

type Count int

// Count type satisfies the stringer interface
func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Log takes stringer as input. It can take any object 
// type which satisfies the stringer interface  
func Log(s fmt.Stringer) {
    log.Println(s.String())
}

func main() {

    c := onlineCourse{"Go Level #1", "Go Programming language", "AK", "dollar" , 199.00 }
    Log(c)

    count := Count(3)
    Log(count)
}
