package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	var left, right []int
	rc := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Invalid line: %s\n", line)
			return
		}
		l, err1 := strconv.Atoi(parts[0])
		r, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing integers: %s\n", line)
			return
		}
		left = append(left, l)
		right = append(right, r)
		rc[r]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	sort.Ints(left)
	sort.Ints(right)
	p1 := 0
	for i := 0; i < len(left); i++ {
		p1 += abs(right[i] - left[i])
	}
	fmt.Println(p1)

	p2 := 0
	for _, l := range left {
		p2 += l * rc[l]
	}
	fmt.Println(p2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

