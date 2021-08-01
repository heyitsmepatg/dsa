package rotateimage

// Explanation of Solution:
// For any nxn matrix, the transpose of that matrix can be determined by
// flipping each element along the diagonal axis from left to right, like so:

// 1 2 3        1 4 7
// 4 5 6  --- > 3 5 8
// 7 8 9        3 6 9

// This transformation can be represented mathetmatically as swap(i, j) for each
// element i,j in matrix M, where swap(i,j) = { matrix(i,j) = matrix(j,i)}

// The diagonal elements remain the same. To get the full 90 degree rotation,
// we must also flip the matrix along the vertical axis. This transformation
// can be represented as flip(i,j) = {matrix(i,j) = matrix(i, n-j-1)}, where
// n is the dimension of the matrix

func rotateImage(matrix [][]int) {
	n := len(matrix)
	//     Transpose Matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	//     Flip horizontally
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = tmp
		}
	}

}
