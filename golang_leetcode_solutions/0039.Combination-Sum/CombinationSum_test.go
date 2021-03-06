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

func TestSearchInsert(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{2, 3, 6, 7},
				7,
			},
			ans{
				[][]int{
					[]int{7},
					[]int{2, 2, 3},
				},
			},
		},
		question{
			para{
				[]int{2, 3, 5},
				8,
			},
			ans{
				[][]int{
					[]int{2, 2, 2, 2},
					[]int{2, 3, 3},
					[]int{3, 5},
				},
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, combinationSum(p.nums, p.target))
	}

}
