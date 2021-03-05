package leetcode

// Given an array nums and a value val, remove all instances of that value in-place and return the new length.

// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

// The order of elements can be changed. It doesn’t matter what you leave beyond the new length.

// Given nums = [3,2,2,3], val = 3,

// Your function should return length = 2, with the first two elements of nums being 2.

// It doesn't matter what you leave beyond the returned length.

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 双指针，i,j分别指向不和val值相等的数组的最后一个元素的下一个下标，j向后遍历
	var i, j = 0, 0
	for ; j < n; j++ {
		// 当前元素和val不相等
		if nums[j] != val {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			i++
		}
	}
	return i
}
