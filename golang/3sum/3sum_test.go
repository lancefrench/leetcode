package problem0015

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
	ans [][]int
}

func TestProblem0015(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{-1, 0, 1, 2, -1, -4},
			},
			output{
				[][]int{
					{-1, 1, 0},
					{-1, 2, -1},
				},
			},
		},
		{
			params{
				[]int{},
			},
			output{
				[][]int{},
			},
		},
		{
			params{
				[]int{0},
			},
			output{
				[][]int{},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		expected := o.ans

		assert.Equal(t, expected, threeSum(p.nums))
	}

}
