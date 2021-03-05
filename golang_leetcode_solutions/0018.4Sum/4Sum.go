package leetcode

import (
	"sort"
)

// Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

// Note:

// The solution set must not contain duplicate quadruplets.

// Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

// A solution set is:
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var n = len(nums)
	for i := 0; i < n-3; {
		for j := i + 1; j < n-2; {
			var start, end = j + 1, n - 1
			for start < end {
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					end--
					for start < end && nums[start] == nums[start-1] {
						start++
					}
					// 不用判断 end + 1 < n, 因为前面有了一次end--
					for start < end && nums[end] == nums[end+1] {
						end--
					}
				} else if sum > target {
					end--
				} else {
					start++
				}
			}
			j++
			for j < n-2 && nums[j] == nums[j-1] {
				j++
			}
		}
		i++
		for i < n-3 && nums[i] == nums[i-1] {
			i++
		}
	}
	return ans
}
