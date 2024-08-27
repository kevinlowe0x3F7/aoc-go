package day4

import (
	"2023/shared"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers []int
	CardNumbers    []int
}

func Day4() {
	lineIter, err := shared.FileLineIterator("./day4/day4.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	var sum int
	var copies [199]int
	for i := 1; i < len(copies); i++ {
		copies[i] = 1
	}
	for line := range lineIter {
		card, err := parseLine(line)
		if err != nil {
			log.Fatal("Failed to parse line", err)
		}
		wins := winningNumbers(card)

		copyCount := copies[card.Id]
		for id := card.Id + 1; id <= card.Id+wins && id < 199; id++ {
			copies[id] += copyCount
		}

		if wins > 0 {
			sum += int(math.Pow(2, float64(wins-1)))
		}
	}
	fmt.Println("Finished calculation part 1")
	fmt.Println(sum)

	var copiesSum int
	for _, copyCount := range copies {
		copiesSum += copyCount
	}

	fmt.Println("Finished calculation part 2")
	fmt.Println(copies)
	fmt.Println(copiesSum)
}

func winningNumbers(card Card) int {
	winningNumbers := 0
	for _, winningNumber := range card.WinningNumbers {
		if slices.Contains(card.CardNumbers, winningNumber) {
			winningNumbers++
		}
	}

	return winningNumbers
}

// parseLine parses a line from the file and returns a CardInfo struct
func parseLine(line string) (Card, error) {
	var info Card

	// Split the line into the part before and after the divider "|"
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return info, fmt.Errorf("invalid format, expected exactly one divider '|'")
	}

	// Parse the card index
	indexPart := strings.Split(parts[0], ":")
	if len(indexPart) != 2 {
		return info, fmt.Errorf("invalid format, expected 'Card X:' format")
	}
	cardIndexStr := strings.TrimSpace(indexPart[0][len("Card"):])
	cardIndex, err := strconv.Atoi(cardIndexStr)
	if err != nil {
		return info, fmt.Errorf("invalid card index: %v", err)
	}
	info.Id = cardIndex

	// Parse the numbers before the divider
	info.WinningNumbers, err = parseNumbers(strings.TrimSpace(indexPart[1]))
	if err != nil {
		return info, fmt.Errorf("invalid numbers before divider: %v", err)
	}

	// Parse the numbers after the divider
	info.CardNumbers, err = parseNumbers(strings.TrimSpace(parts[1]))
	if err != nil {
		return info, fmt.Errorf("invalid numbers after divider: %v", err)
	}

	return info, nil
}

// parseNumbers converts a space-separated string of numbers into a slice of integers
func parseNumbers(numbersStr string) ([]int, error) {
	var numbers []int
	numberParts := strings.Fields(numbersStr)
	for _, part := range numberParts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s': %v", part, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
