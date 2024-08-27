package shared

import (
	"bufio"
	"log"
	"os"
)

func FileLineIterator(filename string) (<-chan string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	lines := make(chan string)

	go func() {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return lines, nil
}
