package packunpackstring

import "testing"

func TestPackString(t *testing.T) {
	// Arrange
	tastTable := []struct {
		str    string
		resStr string
	}{
		{
			str:    "abcdefgh",
			resStr: "1a1b1c1d1e1f1g1h",
		},
		{
			str:    "abbcccddddeeeeeffffffggggggghhhhhhhh",
			resStr: "1a2b3c4d5e6f7g8h",
		},
		{
			str:    "rgb43f",
			resStr: "1r1g1b14131f",
		},
	}

	for _, testCase := range tastTable {
		result := PackString(testCase.str)

		t.Logf("Calling PackString(%v), result %v\n", testCase.str, result)
		if result != testCase.resStr {
			t.Errorf("Incorrect result. Expected %v, got %v", testCase.resStr, result)
		} else {
			t.Logf("Correct result. Expected %v, got %v", testCase.resStr, result)
		}
	}

	// Act

	// Assert
}
