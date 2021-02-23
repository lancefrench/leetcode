package problem0240

import "testing"

type tc struct {
	params
	output
}

type params struct {
	matrix [][]int
	target int
}

type output struct {
	ans bool
}

func Test_Problem0240(t *testing.T) {
	tcs := []tc{
		{
			params{
				[][]int{
					[]int{1, 4, 7, 11, 15},
					[]int{2, 5, 8, 12, 19},
					[]int{3, 6, 9, 16, 22},
					[]int{10, 13, 14, 17, 24},
					[]int{18, 21, 23, 26, 30},
				},
				5,
			},
			output{true},
		},
		{
			params{
				[][]int{
					[]int{1, 4, 7, 11, 15},
					[]int{2, 5, 8, 12, 19},
					[]int{3, 6, 9, 16, 22},
					[]int{10, 13, 14, 17, 24},
					[]int{18, 21, 23, 26, 30},
				},
				20,
			},
			output{false},
		},
	}

	for _, tc := range tcs {
		o, p := tc.output, tc.params

		got := searchMatrix(p.matrix, p.target)
		want := o.ans

		if got != want {
			t.Errorf("got searchMatrix(\"%v\") = %v want %v", tc.params, got, want)
		}
	}
}
