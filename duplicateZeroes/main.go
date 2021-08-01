package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3245/

func main() {

	output := duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
	fmt.Printf("Actual Output: %v\n", output)
	fmt.Printf("Expected Output: %v\n", []int{1, 0, 0, 2, 3, 0, 0, 4})
}

func duplicateZeros(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			shiftArr(arr, i+1)
			arr[i] = 0
			i++
		}
	}
	return arr
}

func shiftArr(arr []int, index int) {
	for j := len(arr) - 1; j >= index; j-- {
		arr[j] = arr[j-1]
	}
}
