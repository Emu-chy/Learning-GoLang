package main

import "fmt"

//why functon need --> for understand an example with a little application

func printWelcomeMessage() {
	fmt.Println("Hi! welcome to the my application")
}

func getUserData() (string, string) {
	//insert user data
	var name string
	fmt.Println("Enter your name...")
	fmt.Scanln(&name)

	var email string
	fmt.Println("Enter your email...")
	fmt.Scanln(&email)

	return name, email

}

func inputCalculateData() (int, int) {
	//do some colculation
	var num1 int
	var num2 int
	fmt.Println("enter the first number")
	fmt.Scanln(&num1)
	fmt.Println("enter the second number")
	fmt.Scanln(&num2)
	return num1, num2
}

// calculation
func Add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

//Dispaly
func Dispaly(name string, email string, sum int) {
	fmt.Println("Hello ", name)
	fmt.Println("your mail addess is: ", email)
	fmt.Println("those number result is: ", sum)
}

// good bye message
func goodBye() {
	// goodbye message
	fmt.Println("Thank you for using application")
	fmt.Println("Good Bye")
}

func main() {
	printWelcomeMessage()
	name, email := getUserData()
	num1, num2 := inputCalculateData()
	sum := Add(num1, num2)
	Dispaly(name, email, sum)
	goodBye()

}
