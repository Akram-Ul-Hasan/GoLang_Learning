package main

import "fmt"

func main() {
	name := "akram"
	if name == "akram" {
		fmt.Println("Welcome Akram!!!")
	}

	number := 0
	if number < 0 {
		fmt.Println("Negative number.")
	} else if number > 0 {
		fmt.Println("Positive number.")
	} else {
		fmt.Println("It's zero.")
	}

	if name == "akram" || number > 0 {
		fmt.Println("Hello")
	}

	if num := 9; num > 0 {
		fmt.Println("Can assign a new variable and use as condition like this")
	}

	fmt.Println("There is no ternary if in Go")
}
