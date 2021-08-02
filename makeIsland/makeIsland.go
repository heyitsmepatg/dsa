package makeIsland

// The main problem is to toggle a 0 to 1 to return the
// largest possible island from a matrix. However, a sub-problem
// that is easier to solve is to determine the existing
// largest island. Once we determine that, we can toggle any
// 0 connected to this graph to a 1, and then return the final
// size of that island

func largestIsland(islands [][]int) int {
	// If islands len == 1 and value is 0, return 1
	if len(islands) == 1 && islands[0][0] == 0 {
		return 1
	}
	maxIslandSize := 0
	// haveVisited := make([][]int, len(islands))

	haveVisited := make([][]int, len(islands))
	for idx := range haveVisited {
		haveVisited[idx] = make([]int, len(islands))
	}

	// Idx for max island
	maxI := 0
	maxJ := 0

	count := 0
	for i, row := range islands {
		for j, v := range row {
			if v == 1 && haveVisited[i][j] != 1 {
				count = 1 // reset count
				buildIsland(islands, haveVisited, i, j, &count)
				if count > maxIslandSize {
					maxIslandSize = count
					maxI = i
					maxJ = j
				}
			}
		}
	}
	return augmentLargestIsland(islands, maxI, maxJ, maxIslandSize)
}

func augmentLargestIsland(islands [][]int, maxI int, maxJ int, maxSize int) int {
	haveVisited := make([][]int, len(islands))
	for idx := range haveVisited {
		haveVisited[idx] = make([]int, len(islands))
	}

	// Construct two slices to represent 4 possible adjacent slice indexes
	neighborRowIdx := []int{-1, 0, 0, 1}
	neighborColIdx := []int{0, -1, 1, 0}

	for k := 0; k < 4; k++ {
		if isSafe(islands, maxI+neighborRowIdx[k], maxJ+neighborColIdx[k], haveVisited) {
			if islands[maxI+neighborRowIdx[k]][maxJ+neighborColIdx[k]] == 0 {
				islands[maxI+neighborRowIdx[k]][maxJ+neighborColIdx[k]] = 1
				maxSize++
				return maxSize
			}
		}
	}
	return maxSize
}

// This is a Depth First Search that constructs the island
func buildIsland(islands [][]int, haveVisited [][]int, i int, j int, count *int) {
	// Mark this cell as visited
	haveVisited[i][j] = 1
	// Default island size

	// Construct two slices to represent 8 possible adjacent slice indexes
	neighborRowIdx := []int{-1, 0, 0, 1}
	neighborColIdx := []int{0, -1, 1, 0}

	for k := 0; k < 4; k++ {
		if isSafe(islands, i+neighborRowIdx[k], j+neighborColIdx[k], haveVisited) {
			if islands[i+neighborRowIdx[k]][j+neighborColIdx[k]] == 1 {
				*count++
				buildIsland(islands, haveVisited, i+neighborRowIdx[k], j+neighborColIdx[k], count)
			}
		}
	}
}

// Returns if an invocation is safe
func isSafe(islands [][]int, row int, col int, haveVisited [][]int) bool {
	return row >= 0 &&
		row < len(haveVisited) &&
		col >= 0 &&
		col < len(haveVisited) &&
		haveVisited[row][col] != 1

}
