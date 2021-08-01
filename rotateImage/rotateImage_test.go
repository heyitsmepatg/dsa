package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotateImage_Example1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expectedMatrix := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	rotateImage(matrix)

	if reflect.DeepEqual(expectedMatrix, matrix) != true {
		t.Error("Expected matrix does not equal actual")
	}
}

func TestRotateImage_Example2(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	expectedMatrix := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	rotateImage(matrix)

	if reflect.DeepEqual(expectedMatrix, matrix) != true {
		t.Error("Expected matrix does not equal actual")
	}
}

func TestRotateImage_Example3(t *testing.T) {
	matrix := [][]int{
		{1},
	}
	expectedMatrix := [][]int{
		{1},
	}
	rotateImage(matrix)

	if reflect.DeepEqual(expectedMatrix, matrix) != true {
		t.Error("Expected matrix does not equal actual")
	}
}

func TestRotateImage_Example4(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	expectedMatrix := [][]int{
		{3, 1},
		{4, 2},
	}
	rotateImage(matrix)

	if reflect.DeepEqual(expectedMatrix, matrix) != true {
		t.Error("Expected matrix does not equal actual")
	}
}
