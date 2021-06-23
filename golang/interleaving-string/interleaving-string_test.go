package problem0097

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	s1 string
	s2 string
	s3 string
}

type output struct {
	ans bool
}

func Test_Problem0097(t *testing.T) {
	tcs := []tc{
		{
			params{
				"aabcc",
				"dbbca",
				"aadbbcbcac",
			},
			output{true},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		expected := o.ans
		actual := isInterleave(p.s1, p.s2, p.s3)
		assert.Equal(t, expected, actual)
	}
}
