package subarray

import "testing"

func TestCheckSubArray(t *testing.T) {
	var increasingSubArray = []int{1, 2, 3, 4, 5}
	var decreasingSubArray = []int{5, 4, 3, 2, 1}

	if !CheckSubArray(true, increasingSubArray) {
		t.Error("Got false, expected true")
	}

	if !CheckSubArray(true, decreasingSubArray) {
		t.Error("Got false, expected true")
	}

	var increasingSubArrayWithOneDecrease = []int{1, 2, 3, 2, 4, 5}
	var decreasingSubArrayWithOneIncrease = []int{5, 4, 3, 4, 2, 1}

	if CheckSubArray(true, increasingSubArrayWithOneDecrease) {
		t.Error("Expected 1, got more")
	}

	if CheckSubArray(true, decreasingSubArrayWithOneIncrease) {
		t.Error("Expected 1, got nil")
	}

	var increasingSubArrayWithTooLowIncrease = []int{1, 2, 3, 3, 6}
	var increasingSubArrayWithTooHighIncrease = []int{1, 2, 3, 4, 8}

	if CheckSubArray(true, increasingSubArrayWithTooLowIncrease) {
		t.Error("Expected 1, got more")
	}

	if CheckSubArray(true, increasingSubArrayWithTooHighIncrease) {
		t.Error("Expected 1, got more")
	}

	var testCase = [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	var errorsInTestCases = 0
	for _, subArray := range testCase {
		if CheckSubArray(true, subArray) {
			errorsInTestCases++
		}
	}

	if errorsInTestCases != 2 {
		t.Error("Expected 2, got", errorsInTestCases)
	}
}
