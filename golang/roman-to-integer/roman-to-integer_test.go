package Problem0013

import (
	"testing"
)

type tc struct {
	params
	output
}

type params struct {
	one string
}

type output struct {
	one int
}

func Test_Problem0013(t *testing.T) {
	tcs := []tc{
		{
			params{"III"},
			output{3},
		},
		{
			params{"IV"},
			output{4},
		},
		{
			params{"IX"},
			output{9},
		},
		{
			params{"LVIII"},
			output{58},
		},
		{
			params{"MCMXCIV"},
			output{1994},
		},
	}

	for _, tc := range tcs {
		o, p := tc.output, tc.params

		got := romanToInt(p.one)
		want := o.one

		if got != want {
			t.Errorf("got romanToInt(\"%v\") = %d want %d", p.one, got, want)
		}
	}
}
