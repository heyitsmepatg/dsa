package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/526/deleting-items-from-an-array/3247/

func main() {
	// nums := []int{1, 2, 0, 5, 3, 2, 2, 4}
	// fmt.Printf("Length of arr is: %d\n", removeElement(nums, 2))
	// fmt.Printf("Nums is: %v\n", nums)

	// nums := []int{3, 2, 2, 3}
	// fmt.Printf("Length of arr is: %d\n", removeElement(nums, 3))
	// fmt.Printf("Nums is: %v\n", nums)

	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	// fmt.Printf("Length of arr is: %d\n", removeElement(nums, 2))
	// fmt.Println("Expected nums is: [0,1,4,0,3]")
	// fmt.Printf("Nums is: %v\n", nums)

	nums := []int{3, 3}
	fmt.Printf("Length of arr is: %d\n", removeElement(nums, 3))
	fmt.Println("Expected nums is: [3, 3]")
	fmt.Printf("Nums is: %v\n", nums)

}

func removeElement(nums []int, val int) int {
	length := len(nums)
	pointer := 0
	for _, n := range nums {
		if n == val {
			length--
			continue
		}
		nums[pointer] = n
		pointer++
	}
	return length
}
