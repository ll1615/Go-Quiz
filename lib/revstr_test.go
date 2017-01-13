package lib

import (
	"testing"
)

var testCase = map[string]string{
	"0123456789":         "9876543210",
	"abcdefg":            "gfedcba",
	" abc 1 2 3 ":        " 3 2 1 cba ",
	" @@ 4 语言 楊 たき En !": "! nE きた 楊 言语 4 @@ ",
}

func TestRevstr(t *testing.T) {
	for input, want := range testCase {
		if output := string(Reverse([]byte(input))); output != want {
			t.Errorf("reverse failed.\n%6s:%s\n%6s:%s\n%6s:%s\n",
				"input", input, "want", want, "output", output)
		}
	}
}
