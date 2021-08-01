package main

import "fmt"

//https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/

func main() {
	// nums := []int{10, 2, 5, 3}
	// fmt.Println("Expected Value: true")
	// fmt.Printf("Actual Value: %v\n", checkIfExist(nums))

	// nums := []int{0, 0}
	// fmt.Println("Expected Value: true")
	// fmt.Printf("Actual Value: %v\n", checkIfExist(nums))

	// nums := []int{7, 1, 14, 11}
	// fmt.Println("Expected Value: true")
	// fmt.Printf("Actual Value: %v\n", checkIfExist(nums))

	// nums := []int{3, 1, 7, 11}
	// fmt.Println("Expected Value: false")
	// fmt.Printf("Actual Value: %v\n", checkIfExist(nums))

	nums := []int{-2, 0, 10, -19, 4, 6, -8}
	fmt.Println("Expected Value: false")
	fmt.Printf("Actual Value: %v\n", checkIfExist(nums))

}

func checkIfExist(arr []int) bool {
	for i, a := range arr {
		for j, b := range arr {
			if i == j {
				continue
			}
			if a == 0 && b == 0 {
				return true
			}
			if a != 0 && float32(b)/float32(a) == 2.0 {
				return true
			}
		}
	}
	return false
}

func doubleExists(arr []int) bool {
	return false
}
