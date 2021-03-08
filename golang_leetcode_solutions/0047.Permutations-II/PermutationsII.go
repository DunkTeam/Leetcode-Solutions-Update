package leetcode

import "sort"

// Given a collection of numbers that might contain duplicates, return all possible unique permutations.

// Example:

// Input: [1,1,2]
// Output:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	var tmp []int
	// 切片记录当前下标是否使用过
	used := make([]int, len(nums))
	backtrack(&ans, tmp, &used, nums)
	return ans
}

func backtrack(ans *[][]int, tmp []int, used *[]int, nums []int) {
	// 结束条件
	if len(tmp) == len(nums) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*ans = append(*ans, t)
		return
	}

	// 每一层循环len(nums)次
	for i := 0; i < len(nums); i++ {
		if (*used)[i] != 0 {
			// 如果已经使用过
			continue
		}
		// 剪枝, 注意最后一个条件
		if i > 0 && nums[i] == nums[i-1] && (*used)[i-1] == 0 {
			continue
		}
		(*used)[i] = 1 // 添加标志
		tmp = append(tmp, nums[i])
		backtrack(ans, tmp, used, nums)
		tmp = tmp[:len(tmp)-1]
		(*used)[i] = 0
	}
}
