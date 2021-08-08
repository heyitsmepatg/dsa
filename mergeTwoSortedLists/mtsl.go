package mergetwosortedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// If both lists empty, return empty.
	// If one of the lists empty, return the other.
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l2 == nil && l1 != nil {
		return l1
	}

	var head *ListNode
	// Determine head of Singly Linked List
	if l1.Val <= l2.Val {
		head = &ListNode{
			Val:  l1.Val,
			Next: nil,
		}
		l1 = l1.Next
	} else {
		head = &ListNode{
			Val:  l2.Val,
			Next: nil,
		}
		l2 = l2.Next
	}
	cur := head
	// Construct list from the head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// If either l1 of l2 has remaining values
	for l1 != nil {
		cur.Next = &ListNode{
			Val:  l1.Val,
			Next: nil,
		}
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = &ListNode{
			Val:  l2.Val,
			Next: nil,
		}
		l2 = l2.Next
		cur = cur.Next

	}
	return head
}

// Explanation of Solution:
// Iterate through both lists until one becomes empty. For each iteration, choose the smaller node from the two lists and append to the tail of the new list. Once one list becomes empty, iterate through the remaining items in the other list until all nodes have been visited. Return the combined list
// Time complexity: Each node is visited once. O(n)
// Space complexity: No new nodes are construction, but rather, the nodes are re-arranged into the correct order. O(1)
// V2 is more space efficient than V1 due to not constructing new nodes pointers
func mergeTwoLists_v2(l1 *ListNode, l2 *ListNode) *ListNode {
	// If both lists empty, return empty.
	// If one of the lists empty, return the other.
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l2 == nil && l1 != nil {
		return l1
	}

	var head *ListNode
	// Determine head of Singly Linked List
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	cur := head
	// Construct list from the head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// If either l1 of l2 has remaining values
	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next

	}
	return head
}
