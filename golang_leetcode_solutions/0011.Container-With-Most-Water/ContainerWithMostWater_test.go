package leetcode

import (
	"fmt"
	"testing"
)

type question struct {
	para1
	ans1
}

type para1 struct {
	height []int
}

type ans1 struct {
	maxArea int
}

func TestMaxArea(t *testing.T) {
	qs := []question{
		{
			para1: para1{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			ans1:  ans1{49},
		},
		{

			para1: para1{[]int{1, 1}},
			ans1:  ans1{1},
		},
		{

			para1: para1{[]int{2, 1}},
			ans1:  ans1{1},
		},
		{

			para1: para1{[]int{4, 3, 2, 1, 4}},
			ans1:  ans1{16},
		},
		{

			para1: para1{[]int{1, 2, 1}},
			ans1:  ans1{2},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 11------------------------\n")
	for _, v := range qs {
		para, _ := v.para1, v.ans1
		fmt.Printf("【input】:%v       【output】:%v\n", para, maxArea(para.height))
	}
	fmt.Printf("\n\n\n")
}
