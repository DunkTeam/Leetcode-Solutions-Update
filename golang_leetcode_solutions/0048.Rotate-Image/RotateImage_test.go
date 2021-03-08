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
	nums [][]int
}
type ans struct {
	a [][]int
}

func TestSearchInsert(t *testing.T) {
	qs := []question{
		question{
			para{
				[][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
					[]int{7, 8, 9},
				},
			},
			ans{
				[][]int{
					[]int{7, 4, 1},
					[]int{8, 5, 2},
					[]int{9, 6, 3},
				},
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v \n", p, a)
		rotate(p.nums)
		fmt.Printf("【actual answer】:%v \n", p.nums)
	}

}
