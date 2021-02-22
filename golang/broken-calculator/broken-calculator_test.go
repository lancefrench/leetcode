package Problem0991

import (
	"testing"
)

type tc struct {
	params
	output
}

type params struct {
	X int
	Y int
}

type output struct {
	ans int
}

func Test_Problem0991(t *testing.T) {
	tcs := []tc{
		{
			params{2, 3},
			output{2},
		},
		{
			params{5, 8},
			output{2},
		},
		{
			params{3, 10},
			output{3},
		},
		{
			params{1024, 1},
			output{1023},
		},
	}

	for _, tc := range tcs {
		o, p := tc.output, tc.params

		got := brokenCalc(p.X, p.Y)
		want := o.ans

		if got != want {
			t.Errorf("got brokenCalc(%d, %d) = %d want %d", p.X, p.Y, got, want)
		}
	}
}
