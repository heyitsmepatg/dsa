package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3251/

func main() {
	// nums := []int{0, 3, 2, 1}
	// fmt.Println("Expected output: true")
	// fmt.Printf("Actual output: %v\n", validMountainArray(nums))

	// nums := []int{2, 1}
	// fmt.Println("Expected output: false")
	// fmt.Printf("Actual output: %v\n", validMountainArray(nums))

	// nums := []int{3, 5, 5}
	// fmt.Println("Expected output: false")
	// fmt.Printf("Actual output: %v\n", validMountainArray(nums))

	// nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println("Expected output: false")
	// fmt.Printf("Actual output: %v\n", validMountainArray(nums))

	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Expected output: false")
	fmt.Printf("Actual output: %v\n", validMountainArray(nums))

}

func validMountainArray(A []int) bool {
	l := len(A)
	i := 0

	// walk up
	for i+1 < l && A[i] < A[i+1] {
		i++
	}
	// can't be first or last element
	if i == 0 || i == l-1 {
		return false
	}
	// walk down
	for i+1 < l && A[i] > A[i+1] {
		i++
	}
	return i == l-1
}
