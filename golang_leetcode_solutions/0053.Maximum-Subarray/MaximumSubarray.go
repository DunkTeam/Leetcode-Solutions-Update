package leetcode

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:

// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	// 区间型动态规划 dp[i]表示nums[0---i]之中的和的最大值
	// dp[i] = max(nums[i]+dp[i-1], nums[i]), 当dp[i-1]为负数，取nums[i]
	ans := nums[0]
	// 初始化
	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(dp[i], ans)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
