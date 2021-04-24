package problem0008

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

func TestProblem0008(t *testing.T) {
	tcs := []tc{
		{
			params{"42"},
			output{42},
		},
		{
			params{"   -42"},
			output{-42},
		},
		{
			params{"4193 with words"},
			output{4193},
		},
		{
			params{"words and 987"},
			output{0},
		},
		{
			params{"-91283472332"},
			output{-2147483648},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		got := myAtoi(p.s)
		want := o.ans

		assert.Equal(t, want, got)
	}
}
