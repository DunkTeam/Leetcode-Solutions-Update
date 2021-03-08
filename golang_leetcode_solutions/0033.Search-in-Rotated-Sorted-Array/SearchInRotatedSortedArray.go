package leetcode

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

// You are given a target value to search. If found in the array return its index, otherwise return -1.

// You may assume no duplicate exists in the array.

// Your algorithm’s runtime complexity must be in the order of O(log n).

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4

// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

// 还需琢磨
func search(nums []int, target int) int {
	// 二分法
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[left] {
			// mid 落在数值大的一部分区间里  left ---mid --- pivot --- right
			if nums[left] <= target && nums[mid] > target {
				// 比较target
				// nums[left] --- target --- nums[mid] -- nums[pivot] --- nums[right]
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			// mid 落在数值小的一部分区间里 left ---- pivot -- mid --- right
			if nums[mid] < target && target <= nums[hight] {
				// nums[left] --- nums[pivot] --nums[mid] --- target ---  nums[right]
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] == nums[mid] {
				left++
			}
			if nums[right] == nums[mid] {
				right--
			}
		}
	}
	return -1
}
