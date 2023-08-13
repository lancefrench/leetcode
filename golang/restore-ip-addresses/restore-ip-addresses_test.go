package problem0093

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
	ans []string
}

func TestProblem0001(t *testing.T) {
	tcs := []tc{
		{
			params{
				"25525511135",
			},
			output{
				[]string{"255.255.11.135", "255.255.111.35"},
			},
		},
		{
			params{
				"0000",
			},
			output{
				[]string{"0.0.0.0"},
			},
		},
		{
			params{
				"101023",
			},
			output{
				[]string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		got := restoreIpAddresses(p.s)
		want := o.ans
		assert.Equal(t, got, want)
	}
}
