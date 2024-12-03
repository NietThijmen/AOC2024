package regex

import "testing"

func TestMultiplyFindAllStringSubmatch(t *testing.T) {
	Init()

	var content = "mul(1, 2) mul(3, 4) mul(5, 6)"
	var amount = -1

	var expected = [][]string{
		{"mul(1, 2)", "1", "2"},
		{"mul(3, 4)", "3", "4"},
		{"mul(5, 6)", "5", "6"},
	}

	result := MultiplyFindAllStringSubmatch(content, amount)

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := 0; i < len(result); i++ {
		if len(result[i]) != len(expected[i]) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("Expected %v, got %v", expected, result)
			}
		}
	}

}
