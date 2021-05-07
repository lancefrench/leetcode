package problem0045

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

func TestProblem0045(t *testing.T) {
	tcs := []tc{
		{
			params{[]int{2, 3, 1, 1, 4}},
			output{2},
		},
		{
			params{[]int{2, 3, 0, 1, 4}},
			output{2},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		actual := jump(p.nums)
		expected := o.ans
		assert.Equal(t, expected, actual)
	}
}
