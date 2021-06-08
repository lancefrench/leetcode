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

	assert.Equal(t, expected, actual)
}
