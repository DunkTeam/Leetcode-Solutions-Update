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
	a []int
}

func TestRemoveDuplicates(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{5, 7, 7, 8, 8, 10},
				8,
			},
			ans{
				[]int{3, 4},
			},
		},
		question{
			para{
				[]int{5, 7, 7, 8, 8, 10},
				6,
			},
			ans{
				[]int{-1, -1},
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, searchRange(p.nums, p.target))
	}

}
