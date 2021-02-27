package problem0946

import "testing"

type tc struct {
	params
	output
}

type params struct {
	pushed []int
	popped []int
}

type output struct {
	ans bool
}

func Test_Problem0946(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{1, 2, 3, 4, 5},
				[]int{4, 5, 3, 2, 1},
			},
			output{true},
		},
		{
			params{
				[]int{1, 2, 3, 4, 5},
				[]int{4, 3, 5, 1, 2},
			},
			output{false},
		},
	}

	for _, tc := range tcs {
		o, p := tc.output, tc.params

		got := validateStackSequences(p.pushed, p.popped)
		want := o.ans

		if got != want {
			t.Errorf("got validateStackSequences(%v, %v) = %v want %v", p.pushed, p.popped, got, want)
		}
	}
}
