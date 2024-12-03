package manipulation

import "strings"

func FindWords(input string, lookFor string) []int {
	// Words to search for
	results := make([]int, 0)

	// Iterate over each word
	var indices []int
	start := 0
	for {
		// Find the index of the word starting from `start`
		index := strings.Index(input[start:], lookFor)
		if index == -1 {
			break
		}
		// Calculate the absolute index
		absoluteIndex := start + index
		indices = append(indices, absoluteIndex)
		// Move start past the current match
		start = absoluteIndex + len(lookFor)
	}

	results = append(results, indices...)

	return results
}
