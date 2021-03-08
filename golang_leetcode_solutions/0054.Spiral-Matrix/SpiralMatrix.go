package leetcode

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

// Example 1:

// Input:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]
// Example 2:

// Input:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
// reference: https://leetcode.com/problems/spiral-matrix/discuss/20573/A-concise-C%2B%2B-implementation-based-on-Directions
func spiralOrder(matrix [][]int) []int {
	row, coloumn := len(matrix), len(matrix[0])
	res := []int{}
	if row == 0 || coloumn == 0 {
		return res
	}

	// spiral clockwise order: right -> down -> left -> up
	// direction matrix
	dirs := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}

	// index of direction
	iDir := 0
	ir, ic := 0, -1 // initial position  !!! it is tricky
	// 针对上面例子2 for horizontal movements, the number of shifts follows: [4, 3, 2]. for vertical movements, the number of shifts follows: [2,1,0]
	steps := []int{coloumn, row - 1} // 记录水平方向和垂直方向需要移动的次数
	for steps[iDir%2] != 0 {
		for i := 0; i < steps[iDir%2]; i++ {
			// 计算新坐标
			ir += dirs[iDir][0]
			ic += dirs[iDir][1]
			res = append(res, matrix[ir][ic])
		}

		steps[iDir%2]--
		iDir = (iDir + 1) % 4
	}
	return res
}

func spiralReverseOrder(matrix [][]int) []int {
	row, coloumn := len(matrix), len(matrix[0])
	res := []int{}
	if row == 0 || coloumn == 0 {
		return res
	}

	// spiral counterclockwise/anticlockwise order: down -> right -> up -> left
	// direction matrix
	dirs := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}

	// index of direction
	iDir := 0
	ir, ic := -1, 0 // initial position
	// 针对上面例子2 for horizontal movements, the number of shifts follows: [3,2,1]. for vertical movements, the number of shifts follows: [3,2,1,0]
	steps := []int{coloumn - 1, row} // 记录水平方向和垂直方向需要移动的次数
	for steps[iDir%2] != 0 {
		for i := 0; i < steps[iDir%2]; i++ {
			// 计算新坐标
			ir += dirs[iDir][0]
			ic += dirs[iDir][1]
			res = append(res, matrix[ir][ic])
		}

		steps[iDir%2]--
		iDir = (iDir + 1) % 4
	}
	return res
}
