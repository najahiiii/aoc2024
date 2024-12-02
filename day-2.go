package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isGood(xs []int) bool {
	isSortedAsc := true
	isSortedDesc := true
	for i := 0; i < len(xs)-1; i++ {
		if xs[i] > xs[i+1] {
			isSortedAsc = false
		}
		if xs[i] < xs[i+1] {
			isSortedDesc = false
		}
	}
	validDiff := true
	for i := 0; i < len(xs)-1; i++ {
		diff := abs(xs[i] - xs[i+1])
		if diff < 1 || diff > 3 {
			validDiff = false
			break
		}
	}
	return (isSortedAsc || isSortedDesc) && validDiff
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	infile := os.Args[1]

	file, err := os.Open(infile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var numbers []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Invalid number in line: %s\n", line)
				return
			}
			numbers = append(numbers, num)
		}
		lines = append(lines, numbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	p1 := 0
	p2 := 0
	for _, xs := range lines {
		if isGood(xs) {
			p1++
		}

		good := false
		for j := 0; j < len(xs); j++ {
			newXs := append(xs[:j:j], xs[j+1:]...)
			if isGood(newXs) {
				good = true
				break
			}
		}
		if good {
			p2++
		}
	}

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
