package main

import "fmt"

func main() {
	x := findNumberOfDigits([]int{12, 345, 2, 6, 7896})
	fmt.Println(x)
}

func findNumberOfDigits(nums []int) int {
	numWithEvenDigits := 0
	for _, num := range nums {
		if numDigits(num)%2 == 0 {
			numWithEvenDigits++
		}
	}
	return numWithEvenDigits
}

func numDigits(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}
