package main

import (
	"fmt"
	"strconv"

	. "github.com/fredrikln/advent-of-code-2021/utils"
)

func main() {
	input := ParseFileToStringSlice("input.txt")

	fmt.Println("Part 1:")
	Part1(input)

	fmt.Println("Part 2:")
	Part2(input)
}

func Part1(report []string) {
	gamma, epsilon := CalculateGammaAndEpsilon(report)
	fmt.Println(gamma * epsilon)
}

func Part2(report []string) {
	o2, co2 := CalculateO2AndCO2(report)
	fmt.Println(o2 * co2)
}

func ParseCommonality(report []string) map[int]map[rune]int {
	commonality := make(map[int]map[rune]int)

	for _, line := range report {
		for index, bit := range line {
			if commonality[index] == nil {
				commonality[index] = make(map[rune]int)
				commonality[index]['0'] = 0
				commonality[index]['1'] = 0
			}

			commonality[index][bit]++
		}
	}

	return commonality
}

func FilterByMostCommon(report []string, index int) []string {
	out := make([]string, 0)

	commonality := ParseCommonality(report)
	commonBit := '0'
	if commonality[index]['1'] >= commonality[index]['0'] {
		commonBit = '1'
	}

	for _, line := range report {
		if line[index] == byte(commonBit) {
			out = append(out, line)
		}
	}

	return out
}

func FilterByLeastCommon(report []string, index int) []string {
	out := make([]string, 0)

	commonality := ParseCommonality(report)

	commonBit := '1'
	if commonality[index]['0'] < commonality[index]['1'] {
		commonBit = '0'
	} else if commonality[index]['0'] == commonality[index]['1'] {
		commonBit = '0'
	}

	for _, line := range report {
		if line[index] == byte(commonBit) {
			out = append(out, line)
		}
	}

	return out
}

func CalculateGammaAndEpsilon(report []string) (int, int) {
	commonality := ParseCommonality(report)

	lineLength := len(report[0])

	var gammaString string
	var epsilonString string

	for i := 0; i < lineLength; i++ {
		if commonality[i]['1'] >= commonality[i]['0'] {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)

	return int(gamma), int(epsilon)
}

func CalculateO2AndCO2(report []string) (int, int) {
	lineLength := len(report[0])

	o2Lines := make([]string, len(report))
	copy(o2Lines, report)
	co2Lines := make([]string, len(report))
	copy(co2Lines, report)

	for i := 0; i < lineLength; i++ {
		if len(o2Lines) != 1 {
			o2Lines = FilterByMostCommon(o2Lines, i)
		}
		if len(co2Lines) != 1 {
			co2Lines = FilterByLeastCommon(co2Lines, i)
		}
	}

	o2, _ := strconv.ParseInt(o2Lines[0], 2, 64)
	co2, _ := strconv.ParseInt(co2Lines[0], 2, 64)

	return int(o2), int(co2)
}
