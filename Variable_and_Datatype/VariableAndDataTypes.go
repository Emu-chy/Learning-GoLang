package main

import "fmt"

func main() {
	var firstName string = "Emon"
	var lastName string = "Chowdhury"
	age := 26
	height := 5.7

	fmt.Println(firstName + lastName)
	fmt.Println(age)
	fmt.Println(height)

	// Data types

	fmt.Printf("the data type of Name is %T\n", firstName+lastName)
	fmt.Printf("the data type of age is %T\n", age)
	fmt.Printf("the data type of height is %T\n", height)

}
