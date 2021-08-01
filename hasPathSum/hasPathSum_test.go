package haspathsum

import (
	"testing"
)

// Example: sum of 22

//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \      \
// 7    2      1

func TestHasPathSumWithSum22(t *testing.T) {
	tn1 := TreeNode{1, nil, nil}
	tn4r := TreeNode{4, nil, &tn1}
	tn13 := TreeNode{13, nil, nil}
	tn8 := TreeNode{8, &tn13, &tn4r}

	tn2 := TreeNode{2, nil, nil}
	tn7 := TreeNode{7, nil, nil}
	tn11 := TreeNode{11, &tn7, &tn2}
	tn4 := TreeNode{4, &tn11, nil}
	tn5 := TreeNode{5, &tn4, &tn8}

	result := hasPathSum(&tn5, 22)
	if !result {
		t.Errorf("Expected: true, Actual: %v", result)
	}
}

func TestHasPathSumWithSum0AndEmptyTree(t *testing.T) {
	result := hasPathSum(nil, 0)
	if result {
		t.Errorf("Expected: false, Actual: %v", result)
	}
}

func TestHasPathSumWithSum1AndSmallTree(t *testing.T) {
	tn2 := TreeNode{2, nil, nil}
	tn1 := TreeNode{1, &tn2, nil}
	result := hasPathSum(&tn1, 1)
	if result {
		t.Errorf("Expected: false, Actual: %v", result)
	}
}

// Test Case:
// [1,2,null,3,null,4,null,5]
//     	   1
//        / \
//       2  nil
//      / \
//     3  nil
//    / \
//   4  nil
//  /
// 5

// sum: 6
// expected output: false
func TestHasPathSumWithSum6AndLargeTree(t *testing.T) {
	tn5 := TreeNode{5, nil, nil}
	tn4 := TreeNode{4, &tn5, nil}
	tn3 := TreeNode{3, &tn4, nil}
	tn2 := TreeNode{2, &tn3, nil}
	tn1 := TreeNode{1, &tn2, nil}
	result := hasPathSum(&tn1, 16)
	if result {
		t.Errorf("Expected: false, Actual: %v", result)
	}
}
