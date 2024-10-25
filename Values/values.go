package main

import "fmt"

func main() {
	fmt.Println("string " + "concatenation")
	fmt.Println("integer sum 5+10 = ", 5+10)
	fmt.Println("Float64 division: 7/5 =", 7/5.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	a := "1"
	fmt.Printf("Types of '1' is %T", a)
}
