package problem0349

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	nums1 []int
	nums2 []int
}

type output struct {
	ans []int
}

func TestProblem0349(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{1, 2, 2, 1},
				[]int{2, 2},
			},
			output{
				[]int{2},
			},
		},
		{
			params{
				[]int{4, 9, 5},
				[]int{9, 4, 9, 8, 4},
			},
			output{
				[]int{9, 4},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		expected := o.ans
		actual := intersection(p.nums1, p.nums2)
		assert.Equal(t, expected, actual)
	}
}
