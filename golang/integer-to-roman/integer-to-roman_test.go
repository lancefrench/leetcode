package problem0012

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	num int
}

type output struct {
	ans string
}

func TestProblem0012(t *testing.T) {
	tcs := []tc{
		{
			params{3},
			output{"III"},
		},
		{
			params{4},
			output{"IV"},
		},
		{
			params{9},
			output{"IX"},
		},
		{
			params{58},
			output{"LVIII"},
		},
		{
			params{1994},
			output{"MCMXCIV"},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		expected := o.ans
		actual := intToRoman(p.num)
		assert.Equal(t, expected, actual)
	}

}
