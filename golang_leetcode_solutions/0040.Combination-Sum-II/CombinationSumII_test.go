package leetcode

import (
	"fmt"
	"sort"
	"testing"

	"github.com/smartystreets/goconvey/convey"
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
	convey.Convey("Test CombinationSum2", t, func() {
		qs := []question{
			question{
				para{
					[]int{10, 1, 2, 7, 6, 1, 5},
					8,
				},
				ans{
					[][]int{
						[]int{1, 7},
						[]int{1, 2, 5},
						[]int{2, 6},
						[]int{1, 1, 6},
					},
				},
			},
			question{
				para{
					[]int{2, 5, 2, 1, 2},
					5,
				},
				ans{
					[][]int{
						[]int{1, 2, 2},
						[]int{5},
					},
				},
			},
		}
		for _, s := range qs {
			p, a := s.para, s.ans
			// t := combinationSum2(p.nums, p.target)
			// convey.So(t, ShouldDoubleSliceEqualWithoutOrder, a)
			fmt.Printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, combinationSum2(p.nums, p.target))
		}
	})
}

// type assertion func(actual interface{}, expected ...interface{}) string
// [][]int{[]int{1,2,3}, []int{2,2,2}}
// 如何比较两个二维切片是否相等
func ShouldDoubleSliceEqualWithoutOrder(actual interface{}, expected ...interface{}) string {

}

func checkSliceEqualWithoutOrder(a []int, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}
	// 排序
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
