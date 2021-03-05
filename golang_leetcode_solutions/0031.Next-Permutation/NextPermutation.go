package leetcode

import "math"

// Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

// If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).

// The replacement must be  in place and use only constant extra memory.

// Input: nums = [1,2,3]
// Output: [1,3,2]

// Input: nums = [3,2,1]
// Output: [1,2,3]

// Input: nums = [1,1,5]
// Output: [1,5,1]

// Input: nums = [1]
// Output: [1]

//  [8,9,6,10,7,2] 6,7--> [8,9,7,10,6,2] --> [8,9,7,2,6,10]
func nextPermutation(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	// j: 从右往左第一个不是升序的元素下标
	// var j = n - 1
	// for ; j >= 0; j-- {
	// 	if j != 0 && nums[j] > nums[j-1] {
	// 		j = j - 1
	// 		break
	// 	}
	// }
	var j = n - 2
	// 这样写相对于上面的写法有个好处：不要判断j+1使用出界；而上面的写法需要判断j-1是否出界
	for ; j >= 0; j-- {
		if nums[j] < nums[j+1] {
			break
		}
	}
	// 没有找到j
	if j == -1 {
		reverse(nums, 0, n-1)
		return
	}

	// i: 从j+1开始，第一个最接近nums[j]的元素的下标，尽量靠右
	var k = j + 1
	var i = 0
	diff := math.MaxInt64
	for ; k < n; k++ {
		tmp := nums[k] - nums[j]
		// 注意这里的判断为<=
		if tmp > 0 && tmp <= diff {
			diff = tmp
			i = k
		}
	}
	// 交换下标i和j的元素
	swap(nums, i, j)
	// 数组下标j+1到末尾 逆序
	reverse(nums, j+1, n-1)

}

// 翻转数组, 对撞指针
func reverse(nums []int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
