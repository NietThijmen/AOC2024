package checker

var Lookup = []map[int]string{
	// M - M
	// - A -
	// S - S
	{
		0: "M",
		2: "M",
		4: "A",
		6: "S",
		8: "S",
	},
	// M - S
	// - A -
	// M - S
	{
		0: "M",
		2: "S",
		4: "A",
		6: "M",
		8: "S",
	},
	// S-M
	// - A -
	// S-M
	{
		0: "S",
		2: "M",
		4: "A",
		6: "S",
		8: "M",
	},

	// S - S
	// - A -
	// M - M
	{
		0: "S",
		2: "S",
		4: "A",
		6: "M",
		8: "M",
	},
}

func splitIntoGrids(arr []string) []string {
	var grids []string

	// -2 because the bottom 2 can't really be a 3x3 grid without whitespace
	for y := 0; y < len(arr); y++ {
		for x := 0; x < len(arr[y]); x++ {

			if x+2 >= len(arr[y]) {
				continue
			}

			if y+2 >= len(arr) {
				break
			}

			println("Found X at", x, y)

			var toAppend string
			toAppend = arr[y][x:x+3] + arr[y+1][x:x+3] + arr[y+2][x:x+3]

			grids = append(grids, toAppend)
		}
	}

	return grids
}

func checkGrid(lookup map[int]string, grid string) bool {
	for key, value := range lookup {
		if string(grid[key]) != value {
			return false
		}
	}

	return true
}

func CheckMasInX(arr []string) int {
	grids := splitIntoGrids(arr)
	var output = 0

	for _, grid := range grids {
		for _, lookup := range Lookup {
			if checkGrid(lookup, grid) {
				output++
			}
		}
	}

	return output

}
