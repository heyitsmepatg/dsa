package symmetrictree

import (
	"testing"
)

func TestIsSymmetricWithSymmetricTree(t *testing.T) {
	tn4l := TreeNode{4, nil, nil}
	tn3l := TreeNode{3, nil, nil}
	tn2l := TreeNode{2, &tn3l, &tn4l}

	tn4r := TreeNode{4, nil, nil}
	tn3r := TreeNode{3, nil, nil}
	tn2r := TreeNode{2, &tn4r, &tn3r}
	root := TreeNode{
		Val:   1,
		Left:  &tn2l,
		Right: &tn2r,
	}
	isSymmetric := isSymmetric(&root)
	if !isSymmetric {
		t.Errorf("Tree is symmetric. Expected Result: True. Actual Result %v", isSymmetric)
	}
}

func TestIsSymmetricWithAsymmetricTree(t *testing.T) {
	tn3r := TreeNode{3, nil, nil}
	tn2r := TreeNode{2, nil, &tn3r}
	root := TreeNode{
		Val:   1,
		Left:  &tn3r,
		Right: &tn2r,
	}
	isSymmetric := isSymmetric(&root)
	if isSymmetric {
		t.Errorf("Tree is asymmetric. Expected Result: False. Actual Result %v", isSymmetric)
	}
}

func TestIsSymmetricRecursiveWithSymmetricTree(t *testing.T) {
	tn4l := TreeNode{4, nil, nil}
	tn3l := TreeNode{3, nil, nil}
	tn2l := TreeNode{2, &tn3l, &tn4l}

	tn4r := TreeNode{4, nil, nil}
	tn3r := TreeNode{3, nil, nil}
	tn2r := TreeNode{2, &tn4r, &tn3r}
	root := TreeNode{
		Val:   1,
		Left:  &tn2l,
		Right: &tn2r,
	}
	isSymmetric := isSymmetricRecursive(&root)
	if !isSymmetric {
		t.Errorf("Tree is symmetric. Expected Result: True. Actual Result %v", isSymmetric)
	}
}

func isSymmetricRecursiveWithAsymmetricTree(t *testing.T) {
	tn3r := TreeNode{3, nil, nil}
	tn2r := TreeNode{2, nil, &tn3r}
	root := TreeNode{
		Val:   1,
		Left:  &tn3r,
		Right: &tn2r,
	}
	isSymmetric := isSymmetricRecursive(&root)
	if isSymmetric {
		t.Errorf("Tree is asymmetric. Expected Result: False. Actual Result %v", isSymmetric)
	}
}
