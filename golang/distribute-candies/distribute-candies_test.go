package problem0575

import "testing"

type tc struct {
	params
	output
}

type params struct {
	candyType []int
}

type output struct {
	ans int
}

func Test_Problem0575(t *testing.T) {
	tcs := []tc{
		{
			params{[]int{1, 1, 2, 2, 3, 3}},
			output{3},
		},
	}

	for _, tc := range tcs {
		a, p := tc.output, tc.params

		got := distributeCandies(p.candyType)
		want := a.ans

		if got != want {
			t.Errorf("got distributeCandies(%v) = %d; want %d", p.candyType, got, want)
		}
	}
}
