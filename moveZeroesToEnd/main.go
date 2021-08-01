package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println("Expected Output: [1,3,12,0,0]")
	moveZeroes(nums)
	fmt.Printf("Actual Output: %v\n", nums)

	// nums := []int{0, 0, 1}
	// fmt.Println("Expected Output: [1,0,0]")
	// moveZeroes(nums)
	// fmt.Printf("Actual Output: %v\n", nums)
}

func moveZeroes(nums []int) {
	if len(nums) == 0 || nums == nil {
		return
	}

	// Swap 0 with non-zero
	idx1 := 0
	for idx2 := 0; idx2 < len(nums); idx2++ {
		if nums[idx2] != 0 {
			nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
			idx1++
		}
	}

	// // Move non-zero values to beginning of array
	// for idx2 < len(nums) {
	// 	if nums[idx2] != 0 {
	// 		nums[idx1] = nums[idx2]
	// 		nums[idx2] = 0
	// 		idx1++
	// 	}
	// 	idx2++
	// }
	// // make rest of the array zero valued
	// for idx1 < len(nums) {
	// 	nums[idx1] = 0
	// 	idx1++
	// }

}
