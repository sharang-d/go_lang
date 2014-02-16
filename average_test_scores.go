package main

import "fmt"

func main() {
	test_scores := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	length := len(test_scores)
	for i := length - 1; i >= 0; i-- {
		sum += test_scores[i]
	}
	fmt.Printf("Average of test scores = %f\n", float32(sum)/float32(length))
}
