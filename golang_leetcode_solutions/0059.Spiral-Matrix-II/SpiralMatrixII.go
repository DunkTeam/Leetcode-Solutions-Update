package leetcode

// Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

// Example:

// Input: 3
// Output:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

func generateMatrix(n int) [][]int {
	// spiral clockwise order: right -> down -> left -> up
	// direction matrix
	dirs := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	// 初始化坐标
	ir, ic := 0, -1
	steps := []int{n, n - 1}
	iDir := 0
	num := 1

	for steps[iDir%2] != 0 {
		for i := 0; i < steps[iDir%2]; i++ {
			ir += dirs[iDir][0]
			ic += dirs[iDir][1]
			ans[ir][ic] = num
			num++
		}
		steps[iDir%2]--
		iDir = (iDir + 1) % 4
	}
	return ans
}
