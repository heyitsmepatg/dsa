package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3253/

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	fmt.Printf("nums1 before merge: %v\n", nums1)
	merge(nums1, m, nums2, n)

	fmt.Printf("nums1 after merge: %v\n", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, n2 := range nums2 {
		nums1[i+m] = n2
	}
	sort.Ints(nums1)
}
