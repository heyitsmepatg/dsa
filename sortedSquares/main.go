package main

import (
	"fmt"
)

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	// fmt.Println("Expected Output: [0,1,9,16,100]")
	// fmt.Printf("Actual Output: %v\n", sortedSquares(nums))

	nums := []int{0, 2}
	fmt.Println("Expected Output: [0,4]")
	fmt.Printf("Actual Output: %v\n", sortedSquares(nums))

}

func sortedSquares(A []int) []int {

	i := 0 // negative array iterator
	j := 0 // positive array iterator

	for j < len(A) && A[j] < 0 {
		j++
	}
	i = j - 1

	B := make([]int, len(A))

	k := 0
	for i >= 0 && j < len(A) {
		if A[i]*A[i] < A[j]*A[j] {
			B[k] = A[i] * A[i]
			i--
		} else {
			B[k] = A[j] * A[j]
			j++
		}
		k++
	}
	for i >= 0 {
		B[k] = A[i] * A[i]
		i--
		k++
	}
	for j < len(A) {
		B[k] = A[j] * A[j]
		j++
		k++
	}
	return B

	// Inefficient, but simple approach

	// for i, a := range A {
	// 	A[i] = a * a
	// }
	// sort.Ints(A)
	// return A
}
