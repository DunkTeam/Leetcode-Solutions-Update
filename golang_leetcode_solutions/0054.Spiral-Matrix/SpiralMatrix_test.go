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
	nums [][]int
}

type ans1 struct {
	a []int
}

func TestMaxArea(t *testing.T) {
	qs := []question{
		{
			para1: para1{[][]int{[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8}, []int{9, 10, 11, 12},
			}},
			ans1: ans1{[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
			// ans1: ans1{[]int{1, 5, 9, 10, 11, 12, 8, 4, 3, 2, 6, 7}},
		},
	}
	for _, v := range qs {
		p, a := v.para1, v.ans1
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, spiralOrder(p.nums))
	}
	fmt.Printf("\n\n\n")
}
