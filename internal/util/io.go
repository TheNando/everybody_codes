package util

import (
	"bufio"
	"os"
	"strings"
)

// Return lines of a file as a slice of strings, trimming whitespace and empty lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
