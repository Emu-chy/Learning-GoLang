package main

import "fmt"

//example of fucntions
// 1. Add and Multiplication two number
func getOutput(a int, b int) (int, int) {
	sum := a + b
	multifly := a * b
	return sum, multifly
}

//example -2 simple string
func universalTruth() {
	fmt.Println("Education must be Free !")
}

func printSomething(name string) {
	fmt.Println("Welcome to the GO Lang Course ", name)

}
func main() {
	a := 10
	b := 20
	result1, result2 := getOutput(a, b)

	fmt.Println(result1, result2)

	universalTruth()

	printSomething("Emon")

}
