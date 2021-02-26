package leetcode

// area = width * height
// 为了让area最大，可以先使得width最大
// 对撞指针
func maxArea(height []int) int {
	var i, j = 0, len(height) - 1
	var ans = 0
	for i < j {
		diff := j - i
		if height[i] < height[j] {
			ans = max(height[i]*diff, ans)
			i++
		} else {
			ans = max(height[j]*diff, ans)
			j--
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
