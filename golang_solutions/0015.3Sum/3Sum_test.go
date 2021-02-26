package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums []int
}

type ans1 struct {
	one [][]int
}

func TestThreeSum(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{-1, 0, 1, 2, -1, -4}},
			ans1{[][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}}},
		},
		{
			para1{[]int{0}},
			ans1{[][]int{}},
		},
		{
			para1{[]int{0, 0, 0}},
			ans1{[][]int{[]int{0, 0, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 15------------------------\n")

	for _, q := range qs {
		ans, p := q.ans1, q.para1
		fmt.Printf("【input】:%v    【expected output】:%v   【actual output】:%v\n", p, ans, threeSum(p.nums))
	}
	fmt.Printf("\n\n\n")
}
