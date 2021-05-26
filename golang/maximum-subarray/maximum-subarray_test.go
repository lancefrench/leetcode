package problem0053

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestProblem0053(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			output{
				6,
			},
		},
		{
			params{
				[]int{1},
			},
			output{
				1,
			},
		},
		{
			params{
				[]int{5, 4, -1, 7, 8},
			},
			output{
				23,
			},
		},
	}

	for _, tc := range tcs {
		o, p := tc.output, tc.params
		actual := maxSubArray(p.nums)
		expected := o.ans

		assert.Equal(t, expected, actual)
	}
}
