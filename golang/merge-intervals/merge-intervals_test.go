package problem0056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	intervals [][]int
}

type output struct {
	ans [][]int
}

func TestProblem0056(t *testing.T) {
	tcs := []tc{
		{
			params{
				[][]int{
					{1, 3}, {2, 6}, {8, 10}, {15, 18},
				},
			},
			output{
				[][]int{
					{1, 6}, {8, 10}, {15, 18},
				},
			},
		},
	}
	for _, tc := range tcs {
		p, o := tc.params, tc.output

		actual := merge(p.intervals)
		expected := o.ans
		assert.Equal(t, expected, actual)
	}
}
