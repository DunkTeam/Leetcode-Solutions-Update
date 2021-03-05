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
	a int
}

func TestThreeSumCloset(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{-1, 2, 1, -4},
				1,
			},
			ans{
				2,
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, threeSumClosest(p.nums, p.target))
	}

}
