package checker

const word = "MAS"

// IsXmasWord just recursively look it up
func IsXmasWord(x int, y int, xInc int, yInc int, depth int, arr []string) bool {
	if depth >= len(word) {
		return true
	}

	if x+xInc < 0 {
		return false
	}
	if y+yInc < 0 {
		return false
	}

	if x+xInc >= len(arr[y]) {
		return false
	}
	if y+yInc >= len(arr) {
		return false
	}
	if word[depth] == arr[y+yInc][x+xInc] {
		return IsXmasWord(x+xInc, y+yInc, xInc, yInc, depth+1, arr)
	}

	return false
}
