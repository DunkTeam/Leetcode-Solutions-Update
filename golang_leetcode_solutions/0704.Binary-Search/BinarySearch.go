package leetcode

// Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		// 当有偶数个元素的时候，需要选择是左中位数还是右中位数
		// 左中位数
		mid := left + (right-left)/2
		// 排除中位数
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 判断是否和target相等
	if nums[left] == target {
		return left
	}
	return -1
}
