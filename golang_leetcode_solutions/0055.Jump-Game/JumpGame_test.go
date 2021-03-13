package leetcode

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	qs := []struct {
		para struct{ nums []int }
		ans  struct{ a bool }
	}{
		{
			para: struct{ nums []int }{[]int{2, 3, 1, 1, 4}},
			ans:  struct{ a bool }{true},
		},
		{
			para: struct{ nums []int }{[]int{3, 2, 1, 0, 4}},
			ans:  struct{ a bool }{false},
		},
	}

	for _, v := range qs {
		p, a := v.para, v.ans
		fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, canJump(p.nums))
	}
	fmt.Printf("\n\n\n")
}
