package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3248/
// https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3258/

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("Len is: %d\n", removeDuplicates(nums))
	fmt.Println("Expected nums is: [0, 1, 2, 3, 4]")
	fmt.Printf("Actual nums is: %v\n", nums)

	// nums := []int{1, 1, 2}
	// fmt.Printf("Len is: %d\n", removeDuplicates(nums))
	// fmt.Println("Expected nums is: [1, 2]")
	// fmt.Printf("Actual nums is: %v\n", nums)

}

// Solution is in-place, meaning no new array space is allocated
// Solution takes O(N) because only a single pass through array is needed
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 || nums == nil {
		return len(nums)
	}
	idx1 := 0
	idx2 := 0

	for idx1 < len(nums) && idx2 < len(nums) {
		if nums[idx1] == nums[idx2] {
			idx2++
		} else {
			idx1++
			nums[idx1] = nums[idx2]
		}
	}

	return idx1 + 1
}
