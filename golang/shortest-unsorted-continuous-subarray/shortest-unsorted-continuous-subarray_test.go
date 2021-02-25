package problem0581

import "testing"

type tc struct {
	params
	output
}

type params struct {
	nums []int
}

type output struct {
	ans int
}

func Test_Problem0581(t *testing.T) {
	tcs := []tc{
		{
			params{[]int{2, 6, 4, 8, 10, 9, 15}},
			output{5},
		},
		{
			params{[]int{1, 2, 3, 4}},
			output{0},
		},
		{
			params{[]int{1}},
			output{0},
		},
	}

	for _, tc := range tcs {
		a, p := tc.output, tc.params

		got := findUnsortedSubarray(p.nums)
		want := a.ans

		if got != want {
			t.Errorf("got findLongestWord(%v) = %d; want %d", p.nums, got, want)
		}
	}
}
