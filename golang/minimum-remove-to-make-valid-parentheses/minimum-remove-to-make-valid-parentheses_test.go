package problem1249

import "testing"

type tc struct {
	params
	output
}

type params struct {
	one string
}

type output struct {
	one string
}

func Test_Problem1249(t *testing.T) {
	tcs := []tc{
		{
			params{"lee(t(c)o)de)"},
			output{"lee(t(c)o)de"},
		},
		{
			params{"a)b(c)d"},
			output{"ab(c)d"},
		},
		{
			params{"))(("},
			output{""},
		},
		{
			params{"(a(b(c)d)"},
			output{"a(b(c)d)"},
		},
	}

	for _, tc := range tcs {
		o, p := tc.output, tc.params

		got := minRemoveToMakeValid(p.one)
		want := o.one

		if got != want {
			t.Errorf("got minRemoveToMakeValid(\"%v\") = \"%v\" want \"%v\"", tc.params, got, want)
		}
	}
}
