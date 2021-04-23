package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	nums   []int
	target int
}

type output struct {
	ans []int
}

func TestProblem0001(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{2, 7, 11, 15},
				9,
			},
			output{
				[]int{0, 1},
			},
		},
		{
			params{
				[]int{3, 2, 4},
				6,
			},
			output{
				[]int{1, 2},
			},
		},
		{
			params{
				[]int{3, 3},
				6,
			},
			output{
				[]int{0, 1},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		got := twoSum(p.nums, p.target)
		want := o.ans
		assert.Equal(t, got, want)
	}
}
