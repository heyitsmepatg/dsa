package main

import (
	"fmt"
	"math"
)

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//    3
//   /  \
//  9   20
//     /  \
//    15   7

// [
//   [3],
//   [9,20],
//   [15,7]
// ]
func main() {
	tn7 := TreeNode{7, nil, nil}
	tn15 := TreeNode{15, nil, nil}
	tn20 := TreeNode{20, &tn15, &tn7}
	tn9 := TreeNode{9, nil, nil}
	tn3 := TreeNode{
		Val:   3,
		Left:  &tn9,
		Right: &tn20,
	}
	fmt.Println("Expected: 3")
	fmt.Printf("Actual: %v\n", maxDepth(&tn3))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1

	//  Method using another function

	// if root == nil {
	// 	return 0
	// }
	// result := 1
	// var recursiveMaxDepth func(root *TreeNode, depth int)
	// recursiveMaxDepth = func(root *TreeNode, depth int) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	if depth > result {
	// 		result = depth
	// 	}
	// 	recursiveMaxDepth(root.Left, depth+1)
	// 	recursiveMaxDepth(root.Right, depth+1)

	// }
	// recursiveMaxDepth(root, result)
	// return result

}
