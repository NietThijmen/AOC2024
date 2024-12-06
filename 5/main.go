package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var rules = [][]int{}
var lines = [][]int{}

func parseInputFile() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			// Rule
			parts := strings.Split(line, "|")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			rules = append(rules, []int{a, b})
			continue
		}

		if strings.Contains(line, ",") {
			// Line
			nums := []int{}
			parts := strings.Split(line, ",")
			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				nums = append(nums, num)
			}
			lines = append(lines, nums)
		}
	}
}

func getRuleForTwoInts(a int, b int) []int {
	for _, rule := range rules {
		if rule[0] == a && rule[1] == b {
			return rule
		}

		if rule[0] == b && rule[1] == a {
			return rule
		}
	}

	return nil
}

func isLineValid(line []int) bool {
	for i := 0; i < len(line)-1; i++ {
		rule := getRuleForTwoInts(line[i], line[i+1])
		if rule == nil {
			continue
		}

		if rule[0] != line[i] || rule[1] != line[i+1] {
			return false
		}
	}
	return true
}

func makeLineValid(line []int) []int {
	var copiedLine = make([]int, len(line))
	copy(copiedLine, line)

	for currentIndex := 0; currentIndex < len(copiedLine)-1; currentIndex++ {
		for otherIndex := 0; otherIndex < len(copiedLine)-1; otherIndex++ {
			if currentIndex == otherIndex {
				continue
			}

			rule := getRuleForTwoInts(copiedLine[currentIndex], copiedLine[otherIndex])
			if rule == nil {
				continue
			}

			if rule[0] == copiedLine[currentIndex] && rule[1] == copiedLine[otherIndex] {
				// Swap
				copiedLine[currentIndex], copiedLine[otherIndex] = copiedLine[otherIndex], copiedLine[currentIndex]
			}
		}
	}

	return copiedLine
}

func getMiddleEntryInLine(line []int) int {
	return line[len(line)/2]
}

func main() {
	parseInputFile()
	// part 1
	var alreadyCorrectLines [][]int
	var answer int = 0
	var answerTwo int = 0
	for _, line := range lines {
		if isLineValid(line) {
			alreadyCorrectLines = append(alreadyCorrectLines, line)
		}
	}

	for _, line := range alreadyCorrectLines {
		answer += getMiddleEntryInLine(line)
	}

	fmt.Printf("Answer: %d\n", answer)

	var correctedIncorrectLines [][]int
	for _, line := range lines {
		if !isLineValid(line) {
			corrected := makeLineValid(line)
			correctedIncorrectLines = append(correctedIncorrectLines, corrected)
		}
	}

	for _, line := range correctedIncorrectLines {
		answerTwo += getMiddleEntryInLine(line)
	}

	fmt.Printf("Answer two: %d\n", answerTwo)

}
