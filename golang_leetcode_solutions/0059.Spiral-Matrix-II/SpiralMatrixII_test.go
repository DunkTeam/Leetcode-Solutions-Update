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
	n int
}

type ans1 struct {
	a [][]int
}

func TestMaxArea(t *testing.T) {
	qs := []question{
		{
			para1{3},
			ans1{[][]int{
				[]int{1, 2, 3},
				[]int{8, 9, 4},
				[]int{7, 6, 5},
			}},
		},
	}
	for _, v := range qs {
		p, a := v.para1, v.ans1
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, generateMatrix(p.n))
	}
	fmt.Printf("\n\n\n")
}
