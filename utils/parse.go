package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func ParseFileToStringSlice(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	output := make([]string, 0)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func ParseFirstRowToIntSlice(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic("Could not read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nums := strings.Split(scanner.Text(), ",")
	numbers := make([]int, 0, len(nums))

	for _, number := range nums {
		n, _ := strconv.Atoi(number)
		numbers = append(numbers, n)
	}

	return numbers
}
