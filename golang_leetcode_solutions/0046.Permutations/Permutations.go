package leetcode

// Given a collection of distinct integers, return all possible permutations.

// Example:

// Input: [1,2,3]
// Output:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

func permute(nums []int) [][]int {
	var ans [][]int
	var tmp []int
	// // map记录当前值是否使用过
	// used := make(map[int]bool)
	// 切片记录当前下标是否使用过
	used := make([]int, len(nums))
	backtrack1(&ans, tmp, &used, nums)
	return ans
}

func backtrack(ans *[][]int, tmp []int, used *map[int]bool, nums []int) {
	// 结束条件
	if len(tmp) == len(nums) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*ans = append(*ans, t)
		return
	}

	// 每一层循环len(nums)次
	for i := 0; i < len(nums); i++ {
		if exist := (*used)[nums[i]]; exist {
			// 如果已经使用过
			continue
		}
		(*used)[nums[i]] = true // 添加标志
		tmp = append(tmp, nums[i])
		backtrack(ans, tmp, used, nums)
		tmp = tmp[:len(tmp)-1]
		(*used)[nums[i]] = false
	}
}

func backtrack1(ans *[][]int, tmp []int, used *[]int, nums []int) {
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
		(*used)[i] = 1 // 添加标志
		tmp = append(tmp, nums[i])
		backtrack1(ans, tmp, used, nums)
		tmp = tmp[:len(tmp)-1]
		(*used)[i] = 0
	}
}
