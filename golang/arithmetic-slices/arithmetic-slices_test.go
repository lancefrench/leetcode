package problem0413

import (
	"testing"
)

type tc struct {
	params
	answer
}

type params struct {
	one []int
}

type answer struct {
	one int
}

func Test_Problem0413(t *testing.T) {
	tcs := []tc{
		{
			params{[]int{1, 2, 3, 4}},
			answer{3},
		},
	}

	for _, tc := range tcs {
		a, p := tc.answer, tc.params

		got := numberOfArithmeticSlices(p.one)
		want := a.one

		if got != want {
			t.Errorf("got numberOfArithmeticSlices(%v) = %d; want %d", tc.params, got, want)
		}
	}
}
