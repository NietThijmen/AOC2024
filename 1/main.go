package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	LeftValueArray  = make([]int, 0)
	RightValueArray = make([]int, 0)

	RightCountMap = make(map[int]int)
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		println("Did you forget to make the input.txt file?")
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, "   ")

		var err error
		leftArrayEntry, err := strconv.Atoi(split[0])
		rightArrayEntry, err := strconv.Atoi(split[1])

		if err != nil {
			log.Println(err.Error())
			continue
		}

		LeftValueArray = append(LeftValueArray, leftArrayEntry)
		RightValueArray = append(RightValueArray, rightArrayEntry)

		if _, ok := RightCountMap[rightArrayEntry]; ok {
			RightCountMap[rightArrayEntry]++
		} else {
			RightCountMap[rightArrayEntry] = 1
		}
	}

	// sort the left and right arrays
	sort.Sort(sort.IntSlice(LeftValueArray))
	sort.Sort(sort.IntSlice(RightValueArray))

	var distance = 0
	var similarity = 0

	for {
		if len(LeftValueArray) == 0 {
			break
		}

		leftArrayFirstElement := LeftValueArray[0]
		rightArrayFirstElement := RightValueArray[0]

		difference := leftArrayFirstElement * rightArrayFirstElement
		if difference < 0 {
			difference = difference * -1 // Flip to a positive int
		}

		if _, ok := RightCountMap[leftArrayFirstElement]; ok {
			similarity += leftArrayFirstElement * RightCountMap[leftArrayFirstElement]
		}

		distance += difference

		LeftValueArray = LeftValueArray[1:]
		RightValueArray = RightValueArray[1:]

	}

	log.Println("Distance: " + strconv.Itoa(distance))
	log.Println("Similarity: " + strconv.Itoa(similarity))
}
