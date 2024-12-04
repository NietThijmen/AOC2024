package main

import (
	"bufio"
	"fmt"
	"github.com/nietthijmen/aoc2024/4/src/checker"
	"log"
	"os"
)

var directions = [8][2]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
}

func main() {
	file, err := os.Open("input.txt")

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	answer := 0
	masInXCount := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'X' {
				for _, direction := range directions {
					xInc := direction[0]
					yInc := direction[1]
					if checker.IsXmasWord(x, y, xInc, yInc, 0, lines) {
						answer++
					}
				}
			}
		}
	}

	masInXCount = checker.CheckMasInX(lines)

	fmt.Printf("Answer: %d\n", answer)
	fmt.Printf("MAS in X count: %d\n", masInXCount)
}
