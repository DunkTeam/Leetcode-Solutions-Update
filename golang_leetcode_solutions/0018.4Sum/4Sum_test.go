package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}
type para struct {
	nums   []int
	target int
}
type ans struct {
	a [][]int
}

func TestFourSum(t *testing.T) {

	qs := []question{
		question{
			para{
				[]int{1, 0, -1, 0, -2, 2},
				0,
			},
			ans{
				[][]int{
					[]int{-1, 0, 0, 1},
					[]int{-2, -1, 1, 2},
					[]int{-2, 0, 0, 2},
				},
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, fourSum(p.nums, p.target))
	}

}
