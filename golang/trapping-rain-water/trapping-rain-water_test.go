package problem0042

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	height []int
}

type output struct {
	ans int
}

func TestProblem0042(t *testing.T) {
	tcs := []tc{
		{
			params{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			output{6},
		},
		{
			params{[]int{4, 2, 0, 3, 2, 5}},
			output{9},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		actual := trap(p.height)
		expected := o.ans
		assert.Equal(t, expected, actual)
	}
}
