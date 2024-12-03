package day1

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/kevinlowe0x3F7/aoc-go/2023/shared"
)

func Day1() {
	lineIter, err := shared.FileLineIterator("./2024/day1/day1.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	counts2 := make(map[int]int)
	for line := range lineIter {
		nums := strings.Fields(line)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)

		prevCount, ok := counts2[num2]
		if !ok {
			counts2[num2] = 1
		} else {
			counts2[num2] = prevCount + 1
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	similarity := 0
	for i, n1 := range list1 {
		n2 := list2[i]
		sum += int(math.Abs(float64(n1 - n2)))

		count, ok := counts2[n1]
		if !ok {
			continue
		} else {
			similarity += n1 * count
		}
	}
	fmt.Println(sum)
	fmt.Println(similarity)
}
