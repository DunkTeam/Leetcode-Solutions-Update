package leetcode

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Note:

// The solution set must not contain duplicate triplets.

// Given array nums = [-1, 0, 1, 2, -1, -4],

// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

import (
	"sort"
)

// GAIN: 去重的逻辑，当前下标为i，计算完之后，i++，之后使用当前i的值和i-1的值比较是否一样
// 还可以使用map进行计数
func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	var ans [][]int

	for i := 0; i < len(nums)-2; {
		sum := -nums[i]
		// 对撞指针
		var start, end = i + 1, len(nums) - 1
		for start < end {
			tmp := nums[start] + nums[end]
			if tmp == sum {
				ans = append(ans, []int{nums[i], nums[start], nums[end]})
				start++
				end--
				// 去重；判断下一个不能和上一个一样，如果一样就跳过
				// 注意判断边界
				for start < end && nums[start] == nums[start-1] {
					start++
				}
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			} else if tmp > sum {
				end--
			} else {
				start++
			}
		}
		i++
		// 判断下一个不能和上一个一样，如果一样就跳过
		for i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}
	}
	return ans
}

func threeSum1(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	var ans [][]int

	for i := 0; i < len(nums)-2; {
		// 对撞指针
		var start, end = i + 1, len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[start], nums[end]})
				start++
				end--
				// 去重；判断下一个不能和上一个一样，如果一样就跳过
				// 注意判断边界
				for start < end && nums[start] == nums[start-1] {
					start++
				}
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
		i++
		// 判断下一个不能和上一个一样，如果一样就跳过
		for i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}
	}
	return ans
}
