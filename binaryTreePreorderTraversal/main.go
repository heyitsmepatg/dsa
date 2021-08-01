package main

import "fmt"

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tn3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	tn2 := TreeNode{
		Val:   2,
		Left:  &tn3,
		Right: nil,
	}

	tn1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &tn2,
	}

	i := preorderTraversal(&tn1)
	fmt.Println("Expected: [1,2,3]")
	fmt.Printf("Actual: %v\n", i)
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var pt func(*TreeNode)
	pt = func(root *TreeNode) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		pt(root.Left)
		pt(root.Right)
	}
	pt(root)
	return result
}
