package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	// golang中不提供set(集合中不存在重复元素)
	// https://studygolang.com/articles/27476?fr=sidebar
	// map记录数组中出现的元素
	// map[int]int  可以使用 []int来代替
	// map[int]struct{} 模拟set
	m := make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	var ans []int
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ans = append(ans, v)
			delete(m, v)
		}
	}
	return ans
}
