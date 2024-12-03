package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/kevinlowe0x3F7/aoc-go/shared"
)

func Day3() {
	lineIter, err := shared.FileLineIterator("./2024/day3/day3.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	sum := 0
	enabled := true
	for line := range lineIter {
		total := 0
		total, enabled = calculateMulSum(line, enabled)
		sum += total
	}

	fmt.Println(sum)
}

func calculateMulSum(input string, enabled bool) (int, bool) {
	// Regular expression to match valid mul(x, y) commands
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(input, -1)

	total := 0
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if strings.HasPrefix(match[0], "mul") && enabled {
			// match[] is the first number (x), match[2] is the second number (y)
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			total += x * y
		}
	}

	fmt.Printf("total: %d\n", total)
	return total, enabled
}
