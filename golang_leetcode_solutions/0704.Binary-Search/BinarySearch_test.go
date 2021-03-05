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

func TestRemoveDuplicates(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{-1, 0, 3, 5, 9, 12},
				9,
			},
			ans{
				4,
			},
		},
		question{
			para{
				[]int{-1, 0, 3, 5, 9, 12},
				2,
			},
			ans{
				-1,
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, search(p.nums, p.target))
	}

}
