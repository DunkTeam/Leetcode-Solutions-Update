package leetcode

// Given an array of non-negative integers, you are initially positioned at the first index of the array.

// Each element in the array represents your maximum jump length at that position.

// Determine if you are able to reach the last index.

// Example 1:

// Input: [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:

// Input: [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum
//              jump length is 0, which makes it impossible to reach the last index.
func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	// Idea is to work backwards from the last index. Keep track of the smallest index that can "jump" to the last index. Check whether the current index can jump to this smallest index.
	// 记录从哪个下标最短能跳到最后
	last := n - 1
	for i := n - 2; i >= 0; i-- {
		// 当前位置能否到跳跃到last
		if nums[i]+i >= last {
			last = i
		}
	}
	return last == 0
}

func canJump1(nums []int) bool {
	n := len(nums)
	if n == 0 || n == 1 {
		return true
	}

	// 记录能跳到的最远距离, 当前在第0个位置，下标为0
	maxLoc := 0

	for i, v := range nums {
		if v > maxLoc {
			return false
		}
		maxLoc = max(maxLoc, i+v)
	}
	return true
}
