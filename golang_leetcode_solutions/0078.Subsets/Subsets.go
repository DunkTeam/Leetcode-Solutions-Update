package leetcode

// Given a set of distinct integers, nums, return all possible subsets (the power set).

// Note: The solution set must not contain duplicate subsets.

// Example:

// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	var tmp []int
	backtrack(&ans, tmp, nums, 0)
	return ans
}

func backtrack(ans *[][]int, tmp []int, nums []int, start int) {
	// 结束条件
	if start == len(nums) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*ans = append(*ans, t)
		return
	}

	// 每一层有两种选择
	// 1.不选下标为start的值
	backtrack(ans, tmp, nums, start+1)
	// 2.选下标为start的值
	tmp = append(tmp, nums[start])
	backtrack(ans, tmp, nums, start+1)
	tmp = tmp[:len(tmp)-1] // 复位
}

func backtrack1(ans *[][]int, tmp []int, nums []int, start int) {
	t := make([]int, len(tmp))
	copy(t, tmp)
	*ans = append(*ans, t)

	// 每一层有 len(nums) - start中选择
	for i := start; i < len(nums); i++ {
		// if i > start && nums[i] == nums[i-1]{
		// 	continue
		// }
		tmp = append(tmp, nums[i])
		backtrack(ans, tmp, nums, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
