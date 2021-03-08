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
	a [][]int
}

func TestSearchInsert(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{1, 1, 2},
			},
			ans{
				[][]int{
					[]int{1, 1, 2},
					[]int{1, 2, 1},
					[]int{2, 1, 1},
				},
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, permuteUnique(p.nums))
	}

}
