package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	s string
}

type output struct {
	ans int
}

func TestProblem0003(t *testing.T) {
	tcs := []tc{
		{
			params{
				"abcabcbb",
			},
			output{
				3,
			},
		},
		{
			params{
				"bbbbb",
			},
			output{
				1,
			},
		},
		{
			params{
				"pwwkew",
			},
			output{
				3,
			},
		},
		{
			params{
				"",
			},
			output{
				0,
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		got := lengthOfLongestSubstring(p.s)
		want := o.ans
		assert.Equal(t, want, got)
	}
}
