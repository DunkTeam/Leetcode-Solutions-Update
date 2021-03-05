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

func TestSearchInser(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{1, 3, 5, 6},
				5,
			},
			ans{
				2,
			},
		},
		question{
			para{
				[]int{1, 3, 5, 6},
				2,
			},
			ans{
				1,
			},
		},
		question{
			para{
				[]int{1, 3, 5, 6},
				7,
			},
			ans{
				4,
			},
		},
		question{
			para{
				[]int{1, 3, 5, 6},
				0,
			},
			ans{
				0,
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, searchInsert(p.nums, p.target))
	}

}
