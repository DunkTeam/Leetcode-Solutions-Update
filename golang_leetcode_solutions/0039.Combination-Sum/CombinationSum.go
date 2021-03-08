package leetcode

import (
	"sort"
)

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

// The same repeated number may be chosen from candidates unlimited number of times.

// Note:

// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:

// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]
// Example 2:

// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

// backtracking question
// https://leetcode.com/problems/combination-sum/discuss/16502/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partitioning)
// 回溯法可能对相同的节点访问多次；深度优先遍历需要记录节点是否被访问过
// 回溯法是建立“状态树”的过程，有可能需要剪枝.  需要思考结束条件；每一层有几种选择
func combinationSum(candidates []int, target int) [][]int {
	// 排序是为了剪枝
	sort.Ints(candidates)

	var ans [][]int
	var tmp []int
	backtrack(&ans, tmp, candidates, target, 0)
	return ans
}

// start表示从数组的第几个开始
// 对切片使用append之后，其底层数组会发生变化, 因此ans必须使用指针
func backtrack(ans *[][]int, tmp []int, candidates []int, remain int, start int) {
	// 结束条件
	if remain < 0 {
		return
	} else if remain == 0 {
		// 必须复制一份
		b := make([]int, len(tmp))
		copy(b, tmp)
		*ans = append(*ans, b)
		return
	}

	// 正常处理逻辑
	// 每一层有len(candidates)-start 个选择
	for i := start; i < len(candidates); i++ {
		// 剪枝
		if candidates[i] > remain {
			break
		}
		tmp = append(tmp, candidates[i])
		backtrack(ans, tmp, candidates, remain-candidates[i], i) // 传入i是因为自身可以被重复选择
		tmp = tmp[:len(tmp)-1]                                   // 删除最后一个, 也就是回溯
	}

}
