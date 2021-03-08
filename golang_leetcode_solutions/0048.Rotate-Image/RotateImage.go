package leetcode

// You are given an n x n 2D matrix representing an image.

// Rotate the image by 90 degrees (clockwise).

// Note:

// You have to rotate the image  in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

// Example 1:

// Given input matrix =
// [
//   [1,2,3],
//   [4,5,6],
//   [7,8,9]
// ],

// rotate the input matrix in-place such that it becomes:
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]
// ]
// Example 2:

// Given input matrix =
// [
//   [ 5, 1, 9,11],
//   [ 2, 4, 8,10],
//   [13, 3, 6, 7],
//   [15,14,12,16]
// ],

// rotate the input matrix in-place such that it becomes:
// [
//   [15,13, 2, 5],
//   [14, 3, 4, 1],
//   [12, 6, 8, 9],
//   [16, 7,10,11]
// ]

func rotate(matrix [][]int) {
	row, column := len(matrix), len(matrix)
	/*
	 * clockwise rotate 顺时针旋转
	 * first reverse up to down, then swap the symmetry
	 * 1 2 3     7 8 9     7 4 1
	 * 4 5 6  => 4 5 6  => 8 5 2
	 * 7 8 9     1 2 3     9 6 3
	 */
	// i代表第几行，j代表第几列
	for i := 0; i < row/2; i++ {
		for j := 0; j < column; j++ {
			matrix[i][j], matrix[row-1-i][j] = matrix[row-1-i][j], matrix[i][j]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
func antiRotate(matrix [][]int) {
	row, column := len(matrix), len(matrix)
	/*
	 * anticlockwise rotate 逆时针旋转
	 * first reverse left to right, then swap the symmetry
	 * 1 2 3     3 2 1     3 6 9
	 * 4 5 6  => 6 5 4  => 2 5 8
	 * 7 8 9     9 8 7     1 4 7
	 */
	// i代表第几行，j代表第几列
	for i := 0; i < row; i++ {
		for j := 0; j < column/2; j++ {
			matrix[i][j], matrix[i][column-1-j] = matrix[i][column-1-j], matrix[i][j]
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
