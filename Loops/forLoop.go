package main

import "fmt"

func main() {
	fmt.Println("There is only for loops in go lang")

	fmt.Println("For as While loop in go lang:")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("Normal For loop in go lang:")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("Infinity loop or always true loop")
	for {
		fmt.Println("loop")
		break
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
