package manipulation

import "testing"

func TestFindWords(t *testing.T) {
	testInput := "mupmuxdo()muxmux,!do()asd"
	lookFor := "do()"

	expected := []int{6, 18}

	result := FindWords(testInput, lookFor)

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}

}
