package leetcode

// Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// Given nums = [1,1,2],

// Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

// It doesn't matter what you leave beyond the returned length.

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	// 双指针 i记录不重复数组的最后一个元素的下标  j往后遍历
	i, j := 0, 1
	for ; j < n; j++ {
		if nums[j] != nums[i] {
			tmp := nums[j]
			nums[j] = nums[i+1]
			nums[i+1] = tmp
			i++
		}
	}

	return i + 1
}
