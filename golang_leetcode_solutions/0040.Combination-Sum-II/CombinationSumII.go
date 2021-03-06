package leetcode

import "sort"

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

// Each number in candidates may only be used once in the combination.

// Note:

// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:

// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// Example 2:

// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ans [][]int
	var tmp []int
	backtrack(&ans, tmp, candidates, target, 0)
	return ans

}
func backtrack(ans *[][]int, tmp []int, candidates []int, remain int, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		// 必须复制一份
		b := make([]int, len(tmp))
		copy(b, tmp)
		*ans = append(*ans, b)
		return
	}

	for i := start; i < len(candidates); i++ {
		// 剪枝
		if candidates[i] > remain {
			break
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		tmp = append(tmp, candidates[i])
		backtrack(ans, tmp, candidates, remain-candidates[i], i+1)
		tmp = tmp[:len(tmp)-1]
	}

}
