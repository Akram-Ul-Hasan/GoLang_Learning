package main

import "fmt"

func main() {
	fmt.Println("Array declaration")

	var arr [5]int
	fmt.Println("emp array: arr = ", arr)

	arr[0] = 108
	fmt.Println("arr[0] =", arr[0])

	fmt.Println("Array Length = ", len(arr))

	fmt.Println("Array declaration(shorthand)")

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array = ", arr1)

	arr1 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", arr1)

	arr = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", arr)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d: ", twoD)
}
