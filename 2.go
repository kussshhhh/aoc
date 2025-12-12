package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(data))
}

// Check if a number is invalid (repeated pattern)
func isInvalid(num int, sz int) bool { // assuming sz is divisible by size
	s := strconv.Itoa(num)
	size := len(s)
	n := size / sz
	if s[0] == '0' {
		return true
	}
	p := s[:sz]
	for i := 0; i < n; i++ {
		p1 := s[i*sz : (i+1)*sz]
		if p != p1 {
			return false
		}
	}

	return true

}

func hehe(num int) bool {
	s := strconv.Itoa(num)
	size := len(s)

	for sz := 1; sz <= size/2; sz++ {
		if size%sz == 0 {
			if size/sz >= 2 && isInvalid(num, sz) {
				return true
			}
		}
	}
	return false
}

func part1(input string) int { // this is also part two sorri
	// Split by commas to get ranges
	ranges := strings.Split(input, ",")

	total := 0
	count := 0
	for _, r := range ranges {
		r = strings.TrimSpace(r)

		parts := strings.Split(r, "-")

		// Convert to integers
		var start, end int
		fmt.Sscanf(parts[0], "%d", &start)
		fmt.Sscanf(parts[1], "%d", &end)

		// Check each number in range
		for num := start; num <= end; num++ {

			if hehe(num) {
				total += num
				count++
			}

		}
	}

	return total
}

func part2(input string) int {

	return 0
}

func main() {
	input := readInput("input2.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
