package leetcode

// Given n non-negative integers a1, a2, …, an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

// Note: You may not slant the container and n is at least 2.

// The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Input: [1,8,6,2,5,4,8,3,7]
// Output: 49

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
