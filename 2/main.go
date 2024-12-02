package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	minimumIncrease = 1
	maximumIncrease = 3

	seperator = " "

	maximumAllowedProblems = 1
)

var (
	safeEntries   = 0
	unsafeEntries = 0

	safeEntriesWithProblemDeduction   = 0
	unsafeEntriesWithProblemDeduction = 0
)

func checkSubArrayForMinMax(a []int) map[int]bool {
	var output = map[int]bool{}
	for id, value := range a {
		if id == len(a)-1 {
			continue
		}

		var diff = value - a[id+1]

		if diff < 0 {
			diff = diff * -1
		}

		if diff > maximumIncrease || diff < minimumIncrease { // diff > 3 or diff < 1
			output[id] = true
		}
	}

	return output
}

func checkSubArrayForAllIncreasingOrDecreasing(a []int) map[int]bool {
	var isIncreasing = true
	var ouput = map[int]bool{}

	var increasingAmount = 0
	var decreasingAmount = 0

	for i := 0; i < 3; i++ {
		if a[i] > a[i+1] {
			decreasingAmount++
		} else {
			increasingAmount++
		}
	}

	if increasingAmount < decreasingAmount {
		isIncreasing = false
	}

	for i := 0; i < 3; i++ {
		if isIncreasing && a[i] > a[i+1] {
			ouput[i] = true
		}

		if !isIncreasing && a[i] < a[i+1] {
			ouput[i] = true
		}
	}

	for id, value := range a {
		if id == len(a)-1 {
			continue
		}

		var diff = value - a[id+1]

		if isIncreasing && diff > 0 {
			ouput[id] = true
		}

		if !isIncreasing && diff < 0 {
			ouput[id] = true
		}
	}

	return ouput
}

func main() {
	// read input.txt to a io.reader
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var intArray = []int{}
		var line = scanner.Text()
		var entries = strings.Split(line, seperator)

		for _, entry := range entries {
			intValue, _ := strconv.Atoi(entry)
			intArray = append(intArray, intValue)
		}

		var problematicIndexes = map[int]bool{}

		// merge the 2 checks into the map
		for k, v := range checkSubArrayForMinMax(intArray) {
			problematicIndexes[k] = v
		}

		for k, v := range checkSubArrayForAllIncreasingOrDecreasing(intArray) {
			problematicIndexes[k] = v
		}

		var problems = len(problematicIndexes)
		if problems > 0 {
			unsafeEntries++
		} else {
			safeEntries++
		}

		if problems > maximumAllowedProblems {
			log.Printf("Unsafe entry: %v", intArray)
			unsafeEntriesWithProblemDeduction++
		} else {
			log.Printf("Safe entry: %v", intArray)
			safeEntriesWithProblemDeduction++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Safe entries: %d", safeEntries)
	log.Printf("Unsafe entries: %d", unsafeEntries)

	log.Printf("Safe entries with problem deduction: %d", safeEntriesWithProblemDeduction)
	log.Printf("Unsafe entries with problem deduction: %d", unsafeEntriesWithProblemDeduction)

}
