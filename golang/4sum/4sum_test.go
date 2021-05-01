package problem0018

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
	ans [][]int
}

func TestProblem0018(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{1, 0, -1, 0, -2, 2},
				0,
			},
			output{
				[][]int{
					{-2, -1, 1, 2},
					{-2, 0, 0, 2},
					{-1, 0, 0, 1},
				},
			},
		},
		{
			params{
				[]int{2, 2, 2, 2, 2},
				8,
			},
			output{
				[][]int{
					{2, 2, 2, 2},
				},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		actual := fourSum(p.nums, p.target)
		expected := o.ans
		assert.Equal(t, expected, actual)
	}
}
