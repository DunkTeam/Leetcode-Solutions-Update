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
	val  int
}
type ans struct {
	a int
}

func TestRemoveDuplicates(t *testing.T) {
	qs := []question{
		question{
			para{
				[]int{3, 2, 2, 3},
				3,
			},
			ans{
				2,
			},
		},
		question{
			para{
				[]int{0, 1, 2, 2, 3, 0, 4, 2},
				2,
			},
			ans{
				5,
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, removeElement(p.nums, p.val))
	}

}
