package problem0623

import "testing"

type tc struct {
	params
	output
}

type params struct {
	root []int
	v    int
	d    int
}

type output struct {
	ans []int
}

func Test_Problem0623(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{4, 2, 3, 1, 6, 5},
				1,
				2,
			},
			output{
				[]int{4, 1, 2, 3, 1, 1, 6, 5},
			},
		},
	}
}
