package main

import (
	"fmt"
	"time"
)

func main() {

		today := time.Now().Weekday()
		fmt.Println("today is - ", today)

		// Simple Switch
		switch time.Now().Weekday() {
		case time.Saturday:
			fmt.Println("Today is Saturday.")
		case time.Sunday:
			fmt.Println("Today is Sunday.")
		default:
			fmt.Println("Today is a weekday.")
		}

		// Case list in a switch case

		switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Today is Weekend.")
		default:
			fmt.Println("Today is a weekday.")
		}

		// Switch case with no value in switch
		// switch init; value
		switch today = time.Now().Weekday() ; {
		case today == time.Saturday || today == time.Sunday:
			fmt.Println("Today is Weekend.")
		default:
			fmt.Println("Today is a weekday.")
		}

		// Switch case with no value in switch
		// switch init; value
		switch today = time.Now().Weekday() ; today {
		case time.Saturday, time.Sunday:
			fmt.Println("Today is Weekend.")
		default:
			fmt.Println("Today is a weekday.")
		}

		// Switch  with multiple true conditions, The first ecnountered truth 
		// will be considered as the final branch to enter
		switch {
		case 1 == 1 :
			fmt.Println("1 == 1")
		case 2 == 2 :
			fmt.Println("2 == 2")
		case 3 == 3 :
			fmt.Println("3 == 3")
		default :
			fmt.Println("Default")
		}
}
