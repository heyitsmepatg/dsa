package symmetrictree

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Example of symmetric tree
//      1
//     / \
//    2   2
//   / \ / \
//  3  4 4  3

func isSymmetric(root *TreeNode) bool {
	// Do level order traversal using queue (BFS)
	// Iterative Approach
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	queue = append(queue, root)

	for len(queue) != 0 {
		node1 := queue[0]
		queue = queue[1:]
		node2 := queue[0]
		queue = queue[1:]

		if node1 == nil && node2 == nil {
			continue
		}
		if (node1 == nil || node2 == nil) || node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left)
		queue = append(queue, node2.Right)
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Left)
	}
	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(tn1 *TreeNode, tn2 *TreeNode) bool {
	// If at this level, all nodes above have been verified
	if tn1 == nil && tn2 == nil {
		return true
	}

	// Covers three conditions that violate symmetry
	if (tn1 == nil && tn2 != nil) || (tn1 != nil && tn2 == nil) || (tn1.Val != tn2.Val) {
		return false
	}

	return isMirror(tn1.Left, tn2.Right) && isMirror(tn1.Right, tn2.Left)
}
