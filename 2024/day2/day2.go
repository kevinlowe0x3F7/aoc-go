package day2

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/kevinlowe0x3F7/aoc-go/shared"
)

func Day2() {
	lineIter, err := shared.FileLineIterator("./2024/day2/day2.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	safe1 := 0
	safe2 := 0

	for line := range lineIter {
		nums := strings.Fields(line)

		if isSafeLine(nums) {
			safe1 += 1
			safe2 += 1
			continue
		}

		for i := range nums {
			removed := removeIndex(nums, i)
			if isSafeLine(removed) {
				safe2 += 1
				break
			}
		}
	}

	fmt.Println(safe1)
	fmt.Println(safe2)
}

func isSafeLine(nums []string) bool {
	var isIncreasing bool
	var n1, n2 int
	foundDirection := false

	// Find the first pair of numbers that are not equal
	for i := 0; i < len(nums)-1; i++ {
		n1, _ = strconv.Atoi(nums[i])
		n2, _ = strconv.Atoi(nums[i+1])
		if n1 == n2 {
			return false // Adjacent equal levels make the report unsafe
		}
		if n1 < n2 {
			isIncreasing = true
			foundDirection = true
			break
		} else if n1 > n2 {
			isIncreasing = false
			foundDirection = true
			break
		}
	}
	if !foundDirection {
		// All numbers are equal, so the report is unsafe
		return false
	}

	// Check the rest of the sequence
	for i := 0; i < len(nums)-1; i++ {
		n1, _ = strconv.Atoi(nums[i])
		n2, _ = strconv.Atoi(nums[i+1])
		if n1 == n2 {
			return false // Adjacent equal levels make the report unsafe
		}
		if isIncreasing && n1 >= n2 {
			return false
		} else if !isIncreasing && n1 <= n2 {
			return false
		}
		diff := int(math.Abs(float64(n1 - n2)))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func removeIndex(s []string, index int) []string {
	result := make([]string, 0, len(s)-1)
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)
	return result
}

// func isSafeLine(nums []string) bool {
// 	num1, _ := strconv.Atoi(nums[0])
// 	num2, _ := strconv.Atoi(nums[1])
//
// 	var isIncreasing bool
// 	if num1 < num2 {
// 		isIncreasing = true
// 	} else {
// 		isIncreasing = false
// 	}
//
// 	for i := 0; i < len(nums)-1; i += 1 {
// 		n1, _ := strconv.Atoi(nums[i])
// 		n2, _ := strconv.Atoi(nums[i+1])
// 		if isIncreasing && n1 >= n2 {
// 			return false
// 		} else if !isIncreasing && n1 <= n2 {
// 			return false
// 		} else {
// 			diff := int(math.Abs(float64(n1 - n2)))
// 			if diff < 1 || diff > 3 {
// 				return false
// 			}
// 		}
// 	}
//
// 	return true
// }
