package main

import (
	"fmt"
	"reflect"
	"github.com/fatih/color"
)

type User struct {
	Name  string `color:"blue"`
	Age   int    `color:"red"`
}

func ColorPrinting(u interface{}) {
	v := reflect.ValueOf(u)
	t := v.Type()

	fmt.Printf("Value - %v , Type - %v \n", v, t)
	fmt.Printf("struct name - %s: \n", t.Name())
	// Iterate over the number of field 
	for i := 0; i < t.NumField(); i++ {
		switch clr := t.Field(i).Tag.Get("color"); clr {
		case "blue":
			color.Set(color.FgBlue)
		case "red":
			color.Set(color.FgRed)
		case "yellow":
			color.Set(color.FgYellow)
		case "green":
			color.Set(color.FgGreen)
		}

		fmt.Printf("%v ", v.Field(i))
		color.Unset()
	}

	fmt.Println("\n-------------------")
}

// To use this program you need to create mod file. 
// As it depends on a color package https://github.com/fatih/color
// go mod init filename.go 
func main() {
	u1 := User{"Ron", 34 }
	u2 := User{"Raj", 46}
	u3 := User{"Bain", 15}

	ColorPrinting(u1)
	ColorPrinting(u2)
	ColorPrinting(u3)
}
