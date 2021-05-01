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
			params{},
			output{},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		actual := fourSum(p.nums, p.target)
		expected := o.ans
		assert.Equal(t, expected, actual)
	}
}
