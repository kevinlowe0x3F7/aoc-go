package day3

import (
	"2023/shared"
	"fmt"
	"log"
	"strconv"
)

type Part struct {
	num    int
	row    int
	minCol int
	maxCol int
}

type Symbol struct {
	symbol rune
	row    int
	col    int
}

type ValidPart struct {
	part   Part
	symbol Symbol
}

func Day3() {
	lineIter, err := shared.FileLineIterator("./day3/day3.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	var grid [140][]rune
	i := 0
	for line := range lineIter {
		grid[i] = []rune(line)
		i++
	}

	var parts []Part
	for row, line := range grid {
		partsFromLine, err := extractNumbers(line, row)
		if err != nil {
			log.Fatal(err)
		}
		parts = append(parts, partsFromLine...)
	}

	var sum int
	var validParts []ValidPart
	for _, part := range parts {
		symbol, ok := isValidPart(part, grid)
		if ok {
			sum += part.num
			validParts = append(validParts, ValidPart{part, symbol})
		}
	}
	fmt.Println("Finished calculating part 1")
	fmt.Println(sum)

	// now need to key symbols by valid parts that surround it
	symbolsToValidParts := make(shared.Multimap[Symbol, ValidPart])
	for _, validPart := range validParts {
		symbolsToValidParts.Put(validPart.symbol, validPart)
	}

	var gearSum int
	for kv := range symbolsToValidParts.Iterator() {
		if kv.Key.symbol == '*' && len(kv.Values) == 2 {
			gearRatio := 1
			for _, part := range kv.Values {
				gearRatio *= part.part.num
			}
			gearSum += gearRatio
		}
	}

	fmt.Println("Finished calculating part 2")
	fmt.Println(gearSum)
}

func extractNumbers(runes []rune, row int) ([]Part, error) {
	var results []Part
	var numberRunes []rune
	minIndex, maxIndex := -1, -1

	// Iterate over the runes
	for i, r := range runes {
		if r >= '0' && r <= '9' { // Check if the rune is a digit
			if minIndex == -1 { // Set the min index when the first digit is found
				minIndex = i
			}
			maxIndex = i // Update the max index as we find more digits
			numberRunes = append(numberRunes, r)
		} else {
			// If we have collected digits and encounter a non-digit, finalize the current number
			if len(numberRunes) > 0 {
				numberStr := string(numberRunes)
				num, err := strconv.Atoi(numberStr)
				if err != nil {
					return nil, err
				}
				results = append(results, Part{
					num,
					row,
					minIndex,
					maxIndex,
				})
				// Reset for the next number
				numberRunes = nil
				minIndex, maxIndex = -1, -1
			}
		}
	}

	// Finalize the last number if the string ended with digits
	if len(numberRunes) > 0 {
		numberStr := string(numberRunes)
		num, err := strconv.Atoi(numberStr)
		if err != nil {
			return nil, err
		}
		results = append(results, Part{
			num,
			row,
			minIndex,
			maxIndex,
		})
	}

	return results, nil
}

func isValidPart(currPart Part, grid [140][]rune) (Symbol, bool) {
	for currRow := currPart.row - 1; currRow <= currPart.row+1; currRow++ {
		for currCol := currPart.minCol - 1; currCol <= currPart.maxCol+1; currCol++ {
			char, ok := safeGet(grid, currRow, currCol)
			if ok && isSymbol(char) {
				return Symbol{char, currRow, currCol}, true
			}
		}
	}
	return Symbol{0, 0, 0}, false
}

func safeGet(grid [140][]rune, row int, col int) (rune, bool) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return '0', false
	}

	return grid[row][col], true
}

func isSymbol(char rune) bool {
	return char != '.' && (char < '0' || char > '9')
}
