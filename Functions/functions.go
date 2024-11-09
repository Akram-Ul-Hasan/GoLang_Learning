package main

import "fmt"

func twoPlus(a int, b int) int {
	return a + b
}
func threePlus(a, b, c int) int {
	return a + b + c
}

// multiple return
func vals() (int, int) {
	return 3, 7
}

// variadic
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("Learning Functions")

	res := twoPlus(5, 4)
	fmt.Println("Result:", res)

	res = threePlus(10, 20, 30)
	fmt.Println("Result:", res)

	// Multiple Return Value
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	// Variadic Functions
	sum(1, 2)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
