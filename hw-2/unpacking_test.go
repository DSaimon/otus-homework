package unpacking

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type TestData struct {
	Str      string
	Expected string
}

func TestUnpackString(t *testing.T) {

	tests := []TestData {
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for _, test := range tests {
		require.Equal(t, test.Expected, unpackString(test.Str), "unpack" + test.Str)
	}
}
