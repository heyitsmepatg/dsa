package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println("Expected Output: [2,4,3,1]")
	fmt.Printf("Actual Output: %v\n", sortArrayByParity(nums))

	// nums := []int{0, 1, 2}
	// fmt.Println("Expected Output: [0,2,1]")
	// fmt.Printf("Actual Output: %v\n", sortArrayByParity(nums))

	// nums := []int{1, 0, 3}
	// fmt.Println("Expected Output: [0,1,3]")
	// fmt.Printf("Actual Output: %v\n", sortArrayByParity(nums))

}

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3260/

func sortArrayByParity(A []int) []int {
	// B := make([]int, len(A))
	// idx1, idx2, idx3 := 0, 0, len(A)-1
	// for idx2 < len(A) {
	// 	if A[idx2]%2 == 0 {
	// 		B[idx1] = A[idx2]
	// 		idx1++
	// 	} else {
	// 		B[idx3] = A[idx2]
	// 		idx3--
	// 	}
	// 	idx2++
	// }
	// return B

	// Attempt to sort in place

	l := len(A)
	idx1 := 0

	for idx2 := 0; idx2 < l; idx2++ {
		if A[idx2]%2 == 0 {
			A[idx2], A[idx1] = A[idx1], A[idx2]
			idx1++
		}

	}
	return A
}
