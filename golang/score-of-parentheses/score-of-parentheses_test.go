package problem0856

import "testing"

type tc struct {
	params
	output
}

type params struct {
	S string
}

type output struct {
	ans int
}

func Test_Problem0856(t *testing.T) {
	tcs := []tc{
		{
			params{"()"},
			output{1},
		},
		{
			params{"(())"},
			output{2},
		},
		{
			params{"()()"},
			output{2},
		},
		{
			params{"(()(()))"},
			output{6},
		},
	}

	for _, tc := range tcs {
		o, p := tc.output, tc.params

		got := scoreOfParentheses(p.S)
		want := o.ans

		if got != want {
			t.Errorf("got scoreOfParantheses(\"%s\") = %d want %d", p.S, got, want)
		}
	}
}
