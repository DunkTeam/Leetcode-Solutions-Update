package leetcode

import (
	"math"
	"sort"
)

// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

// Given array nums = [-1, 2, 1, -4], and target = 1.

// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

// -4 -1 1 2 |  -1
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, diff := 0, math.MaxInt64
	for i := 0; i < len(nums)-2; i++ {
		var start, end = i + 1, len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			// 判断
			if abs(sum-target) < diff {
				res, diff = sum, abs(sum-target)
			}

			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return res

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
