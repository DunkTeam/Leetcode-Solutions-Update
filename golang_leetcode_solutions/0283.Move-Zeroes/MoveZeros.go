package leetcode

// Given an array nums, write a function to move all 0’s to the end of it while maintaining the relative order of the non-zero elements.

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]

// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.
// 和26,27题一样
func moveZeroes(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	var i, j = 0, 0
	for ; j < n; j++ {
		if nums[j] != 0 {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			i++
		}
	}

}
