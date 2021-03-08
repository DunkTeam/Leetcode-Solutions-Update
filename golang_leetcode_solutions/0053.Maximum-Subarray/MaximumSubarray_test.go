package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para1
	ans1
}

type para1 struct {
	nums []int
}

type ans1 struct {
	max int
}

func TestMaxArea(t *testing.T) {
	qs := []question{
		{
			para1: para1{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans1:  ans1{6},
		},
	}
	for _, v := range qs {
		p, a := v.para1, v.ans1
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, maxSubArray(p.nums))
	}
	fmt.Printf("\n\n\n")
}
