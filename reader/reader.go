package aocutils

import (
	"bufio"
	"fmt"
	"os"
)

func Readfile(path string) []string {
	file, err := os.Open(path)
	lines := make([]string, 0)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return lines
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return lines
}
