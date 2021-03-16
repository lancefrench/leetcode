package problem0524

import "testing"

type tc struct {
	params
	output
}

type params struct {
	s string
	d []string
}

type output struct {
	ans string
}

func Test_Problem0524(t *testing.T) {
	tcs := []tc{
		{
			params{
				"abpcplea",
				[]string{"ale", "apple", "monkey", "plea"},
			},
			output{"apple"},
		},
	}

	for _, tc := range tcs {
		a, p := tc.output, tc.params

		got := findLongestWord(p.s, p.d)
		want := a.ans

		if got != want {
			t.Errorf("got findLongestWord(%v) = \"%v\"; want \"%v\"", tc.params, got, want)
		}
	}
}
