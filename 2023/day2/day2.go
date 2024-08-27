package day2

import (
	"2023/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type bag struct {
	red   int
	green int
	blue  int
}

type game struct {
	id    int
	pulls []bag
}

func Day2Part2() {
	lineIter, err := shared.FileLineIterator("./day2/day2.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	var powerSum int
	for line := range lineIter {
		game, err := constructGame(line)
		if err != nil {
			log.Fatal(err)
		}

		powerSum += minimumGame(game)
	}

	fmt.Println("Completed calculation")
	fmt.Println(powerSum)

}

func Day2() {
	lineIter, err := shared.FileLineIterator("./day2/day2.txt")
	if err != nil {
		log.Fatal("Failed to read line by line", err)
		return
	}

	var idSum int
	referenceBag := bag{12, 13, 14}
	for line := range lineIter {
		game, err := constructGame(line)
		if err != nil {
			log.Fatal(err)
		}

		if isValidGame(game, referenceBag) {
			idSum += game.id
		}
	}

	fmt.Println("Completed calculation")
	fmt.Println(idSum)

}

func constructGame(line string) (game, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return game{0, nil}, fmt.Errorf("invalid input format")
	}

	// Extract and convert the game number
	gamePart := strings.TrimSpace(parts[0])
	gameNumberStr := strings.TrimPrefix(gamePart, "Game ")
	gameNumber, err := strconv.Atoi(gameNumberStr)
	if err != nil {
		return game{0, nil}, fmt.Errorf("invalid game number format")
	}

	// The rest of the line contains the bags information
	lineContent := strings.TrimSpace(parts[1])

	// Split the line into segments by semicolon
	segments := strings.Split(lineContent, ";")
	bags := make([]bag, len(segments))

	for i, segment := range segments {
		bags[i] = parseSegment(segment)
	}

	return game{gameNumber, bags}, nil
}

func parseSegment(segment string) bag {
	var bag bag
	items := strings.Split(strings.TrimSpace(segment), ",")

	for _, item := range items {
		parts := strings.Fields(strings.TrimSpace(item))
		if len(parts) != 2 {
			continue
		}
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		color := strings.ToLower(parts[1])

		switch color {
		case "red":
			bag.red = count
		case "blue":
			bag.blue = count
		case "green":
			bag.green = count
		}
	}

	return bag
}

func isValidGame(currGame game, referenceBag bag) bool {
	for _, bag := range currGame.pulls {
		if bag.red > referenceBag.red || bag.blue > referenceBag.blue || bag.green > referenceBag.green {
			return false
		}
	}
	return true
}

// Return the minimum game's bag's red/blue/green values multiplied together
func minimumGame(currGame game) int {
	red := -1
	blue := -1
	green := -1

	for _, pull := range currGame.pulls {
		red = max(red, pull.red)
		blue = max(blue, pull.blue)
		green = max(green, pull.green)
	}

	return red * blue * green
}
