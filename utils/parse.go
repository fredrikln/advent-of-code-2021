package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ParseFileToIntSlice(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic("Could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := make([]int, 0)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Could not parse line")
		}

		output = append(output, value)
	}

	return output
}
