package subarray

import "fmt"

const (
	minimumIncrease = 1
	maximumIncrease = 3
)

// Wrapper function to check the subarray
func CheckSubArray(withOneError bool, sub []int) bool {
	if !withOneError {
		return checkSubArray(sub)
	}

	// check but remove every element once from the subarray
	for i := 0; i < len(sub); i++ {
		subCopy := make([]int, len(sub))
		copy(subCopy, sub)

		if checkSubArray(append(subCopy[:i], subCopy[i+1:]...)) {
			return true
		}
	}

	fmt.Printf("Subarray %v is not safe\n", sub)

	return false
}

// Underlying function to check the subarray
func checkSubArray(sub []int) bool {
	isIncreasing := CheckIfSubArrayIsIncreasing(sub)
	for i := 1; i < len(sub); i++ {
		if isIncreasing && sub[i-1] > sub[i] {
			return false
		}

		if !isIncreasing && sub[i-1] < sub[i] {
			return false
		}

		var diff int

		if isIncreasing {
			diff = sub[i] - sub[i-1]
		} else {
			diff = sub[i-1] - sub[i]
		}

		if diff > maximumIncrease || diff < minimumIncrease {
			return false
		}
	}

	return true
}
