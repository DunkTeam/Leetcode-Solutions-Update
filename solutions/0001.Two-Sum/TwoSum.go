package leetcode

// O(n)
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	// tmp 存储已经遍历过了值以及对应的下标
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := tmp[diff]; ok {
			return []int{v, i}
		}
		// 没有找到
		tmp[nums[i]] = i
	}
	return nil
}

// O(n2)
func twoSum1(nums []int, target int) []int {
	var answer []int
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if diff == nums[j] {
				answer = append(answer, i)
				answer = append(answer, j)
				return answer
			}
		}
	}
	return answer
}
