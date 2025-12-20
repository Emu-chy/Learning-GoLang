package main

import "fmt"

func main() {
	// conditional Statement (==, >, >=, <, <=, &&, ||)
	age := 30
	salary := 30000

	if age == 20 {
		fmt.Println("you are elegiable to merriage")

	} else if age > 20 {
		fmt.Println("you are elegiable to merriage and also ready to go out with your wife :)")
	} else if age >= 20 {
		fmt.Println("you are elegiable to merriage and also ready to take baby :)")
	} else if age == 20 && salary == 30000 {
		fmt.Println("you are elegiable to merriage and also ready to take two baby :)")
	} else if age < 20 || salary == 30000 {
		fmt.Println("you are not elegiable to merriage but if you want, you can love someone :)")
	} else {
		fmt.Println("you are not elegiable to merriage")
	}
}
