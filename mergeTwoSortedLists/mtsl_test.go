package mergetwosortedlists

import (
	"reflect"
	"testing"
)

// Input: l1 = [1,2,4], l2 = [1,3,4]
// Output: [1,1,2,3,4,4]
func TestMergeTwoLists_Example1(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	// [1,1,2,3,4,4]
	expectedOutput := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}

	output := mergeTwoLists_v2(l1, l2)

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Expected List does not match Actual")
	}
}

// Input: l1 = [], l2 = []
// Output: []
func TestMergeTwoLists_Example2(t *testing.T) {
	output := mergeTwoLists_v2(nil, nil)

	if output != nil {
		t.Errorf("Expected: nil. Actual: %v", output)
	}
}

// Input: l1 = [], l2 = [0] Output: [0]
func TestMergeTwoLists_Example3(t *testing.T) {
	output := mergeTwoLists_v2(nil, &ListNode{
		Val:  0,
		Next: nil,
	})

	expectedOutput := &ListNode{
		Val:  0,
		Next: nil,
	}

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("Expected: &{0 <nil>}. Actual: %v", output)
	}
}

// input: [1, 1, 2, 3, 3, 4, 5] [0]
// output: [0, 1, 1, 2, 3, 3, 4, 5]
func TestMergeTwoLists_Example4(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val:  0,
		Next: nil,
	}

	// [1,1,2,3,4,4]
	expectedOutput := &ListNode{
		Val:  0,
		Next: l1,
	}

	output := mergeTwoLists_v2(l1, l2)

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Error("Expected List does not match Actual")
	}
}
