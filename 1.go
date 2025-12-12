package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(data))
}

func hitsDuring(start int, distance int, direction byte) int {
	var k int
	if direction == 'R' {
		k = (100 - start) % 100
	} else {
		k = start % 100
	}

	if k == 0 {
		k = 100
	}

	if distance < k {
		return 0
	}

	return 1 + (distance-k)/100
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	pos := 50
	count := 0

	for _, line := range lines {
		direction := line[0]
		distance := 0
		fmt.Sscanf(line[1:], "%d", &distance)

		rem := distance % 100
		if direction == 'R' {
			pos = (pos + rem) % 100
		} else {
			pos = (pos - rem) % 100
			if pos < 0 {
				pos += 100
			}
		}

		if pos == 0 {
			count++
		}
	}

	return count
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	pos := 50
	count := 0

	for _, line := range lines {
		direction := line[0]
		distance := 0
		fmt.Sscanf(line[1:], "%d", &distance)

		// Simulate each click one by one
		for i := 0; i < distance; i++ {
			if direction == 'R' {
				pos = (pos + 1) % 100
			} else { // 'L'
				pos = (pos - 1 + 100) % 100
			}

			if pos == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	input := readInput("input1.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
