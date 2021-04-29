package problem0016

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
	ans int
}

func TestProblem0016(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{-1, 2, 1, -4},
				1,
			},
			output{2},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		actual := threeSumClosest(p.nums, p.target)
		expected := o.ans
		assert.Equal(t, expected, actual)
	}
}
