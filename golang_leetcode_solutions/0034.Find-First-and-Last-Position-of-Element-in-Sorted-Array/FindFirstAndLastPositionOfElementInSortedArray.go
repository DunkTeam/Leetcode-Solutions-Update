package leetcode

// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

// Your algorithm’s runtime complexity must be in the order of O(log n).

// If the target is not found in the array, return [-1, -1].

// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]

// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

// 二分搜索常见的题目有以下4个
// 代码都是一样的，只是在找到最后的值之后，要判断一下是否将该下标加入到结果中
func searchRange(nums []int, target int) []int {
	// 查找第一个等于目标值的下标
	n := len(nums)
	var ans = []int{-1, -1}
	if n == 0 {
		return ans
	}
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 要找第一个等于目标值的下标
			// 如果mid下标的值小于target，说明mid下标的值肯定不会是目标值
			// 可以排除mid下标
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		ans[0] = left
	}

	// 查找最后一个等于目标值的下标
	left, right = 0, n-1
	for left < right {
		// 使用右中位数
		mid := left + (right-left+1)/2
		if nums[mid] > target {
			// 要找到最后一个等于目标值的下标
			// 如果mid下标的值大于target，说明mid下标的值肯定不会是目标值
			// 可以排除mid下标
			right = mid - 1
		} else {
			left = mid
		}
	}
	if nums[left] == target {
		ans[1] = left
	}
	return ans
}

func search(nums []int, target int) []int {
	// 找到第一个大于或等于目标值的下标
	n := len(nums)
	var ans = []int{-1, -1}
	if n == 0 {
		return ans
	}
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] >= target {
		ans[0] = left
	}
	// 找到最后一个小于或等于目标值的下标
	left, right = 0, n-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if nums[left] <= target {
		ans[1] = left
	}
	return ans
}
