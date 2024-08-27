package day1

import (
	"2023/shared"
	"fmt"
	"log"
	"unicode"
)

func Day1() {
	lineIter, err := shared.FileLineIterator("./day1/day1.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	var sum int
	for line := range lineIter {
		lineNum := getNumFromLine(line)
		sum += lineNum
	}

	fmt.Println("Completed calculation")
	fmt.Println(sum)
}

func getNumFromLine(line string) int {
	runes := []rune(line)
	digits := make([]int, 0)

	for i := 0; i < len(runes); i++ {
		rune := runes[i]
		digit, ok := digitFromLine(rune, i, runes)
		if ok {
			digits = append(digits, digit)
		}

	}

	return digits[0]*10 + digits[len(digits)-1]
}

func digitFromLine(char rune, runeIndex int, runes []rune) (int, bool) {
	if unicode.IsDigit(char) {
		return int(char - '0'), true
	}

	lineLength := len(runes)
	// Check for the start of each digit word
	switch char {
	case 'o':
		if runeIndex+2 < lineLength && runes[runeIndex+1] == 'n' && runes[runeIndex+2] == 'e' {
			return 1, true
		}
	case 't':
		if runeIndex+2 < lineLength {
			if runes[runeIndex+1] == 'w' && runes[runeIndex+2] == 'o' {
				return 2, true
			}
			if runeIndex+4 < lineLength {
				if runes[runeIndex+1] == 'h' && runes[runeIndex+2] == 'r' && runes[runeIndex+3] == 'e' && runes[runeIndex+4] == 'e' {
					return 3, true
				}
			}
		}
	case 'f':
		if runeIndex+3 < lineLength {
			if runes[runeIndex+1] == 'o' && runes[runeIndex+2] == 'u' && runes[runeIndex+3] == 'r' {
				return 4, true
			}
			if runeIndex+3 < lineLength && runes[runeIndex+1] == 'i' && runes[runeIndex+2] == 'v' && runes[runeIndex+3] == 'e' {
				return 5, true
			}
		}
	case 's':
		if runeIndex+2 < lineLength {
			if runes[runeIndex+1] == 'i' && runes[runeIndex+2] == 'x' {
				return 6, true
			}
			if runeIndex+4 < lineLength && runes[runeIndex+1] == 'e' && runes[runeIndex+2] == 'v' && runes[runeIndex+3] == 'e' && runes[runeIndex+4] == 'n' {
				return 7, true
			}
		}
	case 'e':
		if runeIndex+4 < lineLength && runes[runeIndex+1] == 'i' && runes[runeIndex+2] == 'g' && runes[runeIndex+3] == 'h' && runes[runeIndex+4] == 't' {
			return 8, true
		}
	case 'n':
		if runeIndex+3 < lineLength && runes[runeIndex+1] == 'i' && runes[runeIndex+2] == 'n' && runes[runeIndex+3] == 'e' {
			return 9, true
		}
	}
	return 0, false
}
