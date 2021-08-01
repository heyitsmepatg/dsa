package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0}))
}

func findMaxConsecutiveOnes(nums []int) int {
	curMax := 0
	totalMax := 0

	for _, num := range nums {
		if num == 1 {
			curMax++
		} else {
			curMax = 0
		}
		if totalMax < curMax {
			totalMax = curMax
		}
	}
	return totalMax
}
