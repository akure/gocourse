package main

import "fmt"

func main() {

	//Create an empty map using make keyword. 
	// <k,v> => <string, string> 
	m1 := make(map[string]string)
	fmt.Printf("m1 - %v, len = %d \n", m1, len(m1));

	// insert using sub operator
	m1["Rob"] = "Australia"
	m1["Raj"] = "India"
	m1["Ahmed"] = "Iran"
	fmt.Printf("m1 - %v, len = %d \n", m1, len(m1));

	// delete 
	delete(m1, "Rob");
	fmt.Printf("m1 - %v, len = %d \n", m1, len(m1));

	// Value only find
	v := m1["Raj"]
	fmt.Printf("v - %v\n", v)

	v = m1["Rahul"]
	fmt.Printf("v - %v\n", v)

	// value, ok find, 
	// empty place holder ( blank identifier can be used in place 
	// of v  to findout if an item exit or not in the map
	v1, ok := m1["Raj"]
	fmt.Printf("ok - %v, v1 - %v\n",ok, v1)
}
