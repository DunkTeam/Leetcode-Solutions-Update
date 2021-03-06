package leetcode

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

func testRemoveDuplicates(t *testing.t) {
	qs := []question{
		question{
			para{
				[]int{4, 5, 6, 7, 0, 1, 2},
				0,
			},
			ans{
				4,
			},
		},
		question{
			para{
				[]int{4, 5, 6, 7, 0, 1, 2},
				3,
			},
			ans{
				-1,
			},
		},
	}

	for _, s := range qs {
		p, a := s.para, s.ans
		fmt.printf("【input】:%v    【expected answer】:%v  【actual answer】:%v\n", p, a, search(p.nums, p.target))
	}

}
