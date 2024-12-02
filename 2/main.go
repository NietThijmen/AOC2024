package main

import (
	"bufio"
	"github.com/nietthijmen/aoc2024/2/src/subarray"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	separator = " "
)

var (
	safeEntries   = 0
	unsafeEntries = 0

	safeEntriesWithProblemDeduction   = 0
	unsafeEntriesWithProblemDeduction = 0
)

func main() {
	// read input.txt to an io.reader
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var index = 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var intArray []int
		var line = scanner.Text()
		var entries = strings.Split(line, separator)

		for _, entry := range entries {
			intValue, _ := strconv.Atoi(entry)
			intArray = append(intArray, intValue)
		}

		var success = subarray.CheckSubArray(false, intArray)
		if !success {
			unsafeEntries++
		} else {
			safeEntries++
		}

		var successWithProblemDeduction = subarray.CheckSubArray(true, intArray)
		if !successWithProblemDeduction {
			unsafeEntriesWithProblemDeduction++
		} else {
			safeEntriesWithProblemDeduction++
		}

		index++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Safe entries: %d", safeEntries)
	log.Printf("Unsafe entries: %d", unsafeEntries)

	log.Printf("Safe entries with problem deduction: %d", safeEntriesWithProblemDeduction)
	log.Printf("Unsafe entries with problem deduction: %d", unsafeEntriesWithProblemDeduction)

}
