package main

import (
	"fmt"
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
	fmt.Println("Expected: [[3], [9,20], [15, 7]]")
	fmt.Printf("Actual: %v\n", levelOrder(&tn3))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	// result = append(result, []int{root.Val})
	// Attempt 2 - use a queue
	if root == nil {
		return nil
	}
	var result [][]int
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	rowNodes := []int{}
	rowSize := len(queue)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		rowSize--
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		rowNodes = append(rowNodes, node.Val)
		if rowSize == 0 {
			tmp := make([]int, len(rowNodes))
			copy(tmp, rowNodes)
			rowNodes = nil
			result = append(result, tmp)
			rowSize = len(queue)
		}
	}

	// Attempt 1 - incorrect
	// result = append(result, []int{root.Val})

	// var lo func(*TreeNode)
	// lo = func(root *TreeNode) {
	// 	if root == nil {
	// 		return
	// 	}
	// 	var tmp []int
	// 	if root.Left != nil {
	// 		tmp = append(tmp, root.Left.Val)
	// 	}
	// 	if root.Right != nil {
	// 		tmp = append(tmp, root.Right.Val)
	// 	}
	// 	if len(tmp) == 0 {
	// 		return
	// 	}
	// 	result = append(result, tmp)
	// 	lo(root.Left)
	// 	lo(root.Right)

	// }
	// lo(root)

	return result
}
