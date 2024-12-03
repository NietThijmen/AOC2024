package main

import (
	"fmt"
	"github.com/nietthijmen/aoc2024/3/src/manipulation"
	"github.com/nietthijmen/aoc2024/3/src/regex"
	"github.com/nietthijmen/aoc2024/3/src/types"
	"io"
	"os"
	"strconv"
)

func main() {
	var err error
	var answer int = 0
	var answerWithDosAndDonts int = 0

	regex.Init()

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(file)
	contentString := string(content)
	if err != nil {
		panic(err)
	}

	doWordIndices := manipulation.FindWords(contentString, "do()")
	dontWordIndices := manipulation.FindWords(contentString, "don't()")
	var combinedIndices = make([]types.LineStruct, 0)

	combinedIndices = append(combinedIndices, types.LineStruct{0, "do"})

	for _, index := range doWordIndices {
		combinedIndices = append(combinedIndices, types.LineStruct{index, "do"})
	}

	for _, index := range dontWordIndices {
		combinedIndices = append(combinedIndices, types.LineStruct{index, "dont"})
	}

	// order combinedIndices by begin
	for i := 0; i < len(combinedIndices); i++ {
		for j := i + 1; j < len(combinedIndices); j++ {
			if combinedIndices[i].Begin > combinedIndices[j].Begin {
				combinedIndices[i], combinedIndices[j] = combinedIndices[j], combinedIndices[i]
			}
		}
	}

	fmt.Printf("Combined indices: %v\n", combinedIndices)
	fmt.Printf("Length of file: %d\n", len(contentString))

	for id, line := range combinedIndices {
		contentStringCopy := contentString

		var nextLineNum = 0
		if id != len(combinedIndices)-1 {
			var nextLine = combinedIndices[id+1]
			nextLineNum = nextLine.Begin
		} else {
			nextLineNum = len(contentStringCopy) - 1
		}

		var subString = contentStringCopy[line.Begin:nextLineNum]

		matches := regex.MultiplyFindAllStringSubmatch(subString, -1)

		for _, match := range matches {
			if len(match) != 3 {
				continue
			}

			firstOperand := match[1]
			secondOperand := match[2]

			firstOperandInt, err := strconv.Atoi(firstOperand)
			if err != nil {
				panic(err)
			}

			secondOperandInt, err := strconv.Atoi(secondOperand)
			if err != nil {
				panic(err)
			}

			answer += firstOperandInt * secondOperandInt
			if line.Action == "do" {
				answerWithDosAndDonts += firstOperandInt * secondOperandInt
			}
		}
	}

	fmt.Printf("Answer: %d\n", answer)
	fmt.Printf("Answer with do's and dont's: %d\n", answerWithDosAndDonts)
	fmt.Printf("Difference between the 2 answers: %d\n", answer-answerWithDosAndDonts)
}
