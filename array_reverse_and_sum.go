package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for i := len(array) - 1; i >= 0; i-- {
		sum += array[i]
		fmt.Printf("Element %d: %d\n", i, array[i])
	}
	fmt.Printf("Sum of elements = %d\n", sum)
}
