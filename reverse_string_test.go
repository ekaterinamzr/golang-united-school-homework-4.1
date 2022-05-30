package reverse_string

import "testing"

type testPairSlice struct {
	str      string
	reversed string
}

func TestReverseString(t *testing.T) {
	var test testPairSlice = testPairSlice{"Hello world!", "!dlrow olleH"}

	res := ReverseString(test.str)

	if res != test.reversed {
		t.Error("Expected ", test.reversed,
			"got ", res)
	}
}
