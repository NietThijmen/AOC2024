package subarray

func CheckIfSubArrayIsIncreasing(sub []int) bool {
	increasingCount := 0
	decreasingCount := 0

	for i := 0; i < len(sub)-1; i++ {
		if sub[i] > sub[i+1] {
			decreasingCount++
		} else {
			increasingCount++
		}
	}

	return increasingCount > decreasingCount // if true, it is increasing
}
