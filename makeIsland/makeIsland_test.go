package makeIsland

import (
	"testing"
)

func TestLargestIsland_Example1(t *testing.T) {
	grid := [][]int{
		{1, 0},
		{0, 1}}

	largestIslandSize := largestIsland(grid)

	if largestIslandSize != 3 {
		t.Errorf("Expected: 3. Actual: %v", largestIslandSize)
	}
}

func TestLargestIsland_Example2(t *testing.T) {
	grid := [][]int{
		{1, 1},
		{1, 0},
	}

	largestIslandSize := largestIsland(grid)

	if largestIslandSize != 4 {
		t.Errorf("Expected: 4. Actual: %v", largestIslandSize)
	}
}

func TestLargestIsland_Example3(t *testing.T) {
	grid := [][]int{
		{1, 1},
		{1, 1},
	}

	largestIslandSize := largestIsland(grid)

	if largestIslandSize != 4 {
		t.Errorf("Expected: 4. Actual: %v", largestIslandSize)
	}
}

func TestLargestIsland_Example4(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 1},
	}

	largestIslandSize := largestIsland(grid)

	if largestIslandSize != 5 {
		t.Errorf("Expected: 5. Actual: %v", largestIslandSize)
	}
}

// [[0]]

func TestLargestIsland_Example5(t *testing.T) {
	grid := [][]int{{0}}

	largestIslandSize := largestIsland(grid)

	if largestIslandSize != 1 {
		t.Errorf("Expected: 1. Actual: %v", largestIslandSize)
	}

}

// [[1,0,1],[0,0,0],[0,1,1]]
func TestLargestIsland_Example6(t *testing.T) {
	grid := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{0, 1, 1},
	}

	largestIslandSize := largestIsland(grid)

	if largestIslandSize != 4 {
		t.Errorf("Expected: 4. Actual: %v", largestIslandSize)
	}
}
