package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/

func main() {
	nums := []int{17, 18, 5, 4, 6, 1}
	fmt.Println("Expected Result: [18,6,6,6,1,-1]")
	fmt.Printf("Actual Result: %v\n", replaceElements(nums))
}

// Approach starting at end of array, replacing each value with the maxValue.
// Then change the maxValue if a new one occurs
func replaceElements(arr []int) []int {
	maxVal := -1
	for i := len(arr) - 1; i >= 0; i-- {
		a := arr[i]
		arr[i] = maxVal
		if a > maxVal {
			maxVal = a
		}
	}
	return arr
}

// Approach starting at beginning of array. Efficient on memory, super slow
// func replaceElements(arr []int) []int {
// 	if len(arr) == 1 {
// 		arr[0] = -1
// 		return arr
// 	}
// 	curMax := 1
// 	idx1 := 0
// 	idx2 := 0
// 	for idx1 < len(arr) {
// 		for idx2 < len(arr) {
// 			if arr[idx2] > curMax {
// 				curMax = arr[idx2]
// 			}
// 			idx2++
// 		}
// 		arr[idx1] = curMax
// 		curMax = 1
// 		idx1++
// 		idx2 = idx1 + 1
// 	}
// 	arr[len(arr)-1] = -1

// 	return arr
// }
