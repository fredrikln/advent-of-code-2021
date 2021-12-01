package main

import (
	"fmt"

	u "github.com/fredrikln/advent-of-code-2021/utils"
)

func NumIncreasing(depths []int) int {
	last := depths[0]
	count := 0
	for _, depth := range depths[1:] {
		if depth > last {
			count = count + 1
		}
		last = depth
	}

	return count
}

func GetSumSlidingWindow(index int, depths []int) int {
	return depths[index] + depths[index+1] + depths[index+2]
}

func NumIncreasingSlidingWindow(depths []int) int {
	last := GetSumSlidingWindow(0, depths)
	count := 0
	// Loop from index 1 until len-2 (because the last 2 are included in the sliding window)
	for i := 1; i < len(depths)-2; i += 1 {
		total := GetSumSlidingWindow(i, depths)
		if total > last {
			count = count + 1
		}

		last = total
	}

	return count
}

func main() {
	input := u.ParseFileToIntSlice("input.txt")

	fmt.Println("Part 1:")
	Part1(input)
	fmt.Println("Part 2:")
	Part2(input)
}

func Part1(input []int) {
	fmt.Println(NumIncreasing(input))
}

func Part2(input []int) {
	fmt.Println(NumIncreasingSlidingWindow(input))
}
