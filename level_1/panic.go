package main

import (
	"fmt"
)
func fullName(firstName *string, lastName *string) {
    defer fmt.Println("deferred function call - fullname")
    if firstName == nil {
        panic("Panic : firstName can not be nil")
    }
    if lastName == nil {
        panic("Panic: lastName can not be nil")
    }


}

func main() {
    defer fmt.Println("deferred call in main")
    firstName := "Ron"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}

