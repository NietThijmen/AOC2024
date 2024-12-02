package subarray

import "testing"

func TestCheckIfSubArrayIsIncreasing(t *testing.T) {
	var increasingSubArray = []int{1, 2, 3, 4, 5}
	var decreasingSubArray = []int{5, 4, 3, 2, 1}

	if !CheckIfSubArrayIsIncreasing(increasingSubArray) {
		t.Error("Expected true, got false")
	}

	if CheckIfSubArrayIsIncreasing(decreasingSubArray) {
		t.Error("Expected false, got true")
	}

	var increasingSubArrayWithOneDecrease = []int{1, 2, 3, 2, 4, 5}
	var decreasingSubArrayWithOneIncrease = []int{5, 4, 3, 4, 2, 1}

	if !CheckIfSubArrayIsIncreasing(increasingSubArrayWithOneDecrease) {
		t.Error("Expected true, got false")
	}

	if CheckIfSubArrayIsIncreasing(decreasingSubArrayWithOneIncrease) {
		t.Error("Expected false, got true")
	}

}
