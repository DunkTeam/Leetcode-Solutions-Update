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
	nums []int
}
type ans struct {
	a int
}

func TestRemoveDuplicates(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{1, 1, 2},
			},
			ans{
				2,
			},
		},
		question{
			para{
				[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			ans{
				5,
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, removeDuplicates(p.nums))
	}

}
