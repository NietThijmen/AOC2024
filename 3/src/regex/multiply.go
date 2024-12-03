package regex

import "regexp"

var mulRegex = "mul\\((\\d{1,3}),\\s*(\\d{1,3})\\)"
var compiledMulRegex *regexp.Regexp

func Init() {
	compiledMultiplyRegex, err := regexp.Compile(mulRegex)
	if err != nil {
		panic("Failed to compile regex")
	}

	compiledMulRegex = compiledMultiplyRegex
}

func MultiplyFindAllStringSubmatch(content string, amount int) [][]string {
	return compiledMulRegex.FindAllStringSubmatch(content, amount)
}
