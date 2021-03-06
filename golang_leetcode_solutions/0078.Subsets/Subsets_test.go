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
				[]int{1, 2, 3},
			},
			ans{
				[][]int{
					[]int{3},
					[]int{1},
					[]int{2},
					[]int{1, 2, 3},
					[]int{1, 3},
					[]int{2, 3},
					[]int{1, 2},
					[]int{},
				},
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, subsets(p.nums))
	}

}
