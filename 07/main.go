package main

import (
	"fmt"
	"math"

	. "github.com/fredrikln/advent-of-code-2021/utils"
)

func main() {
	numbers := ParseFirstRowToIntSlice("input.txt")

	var maxNum int
	for _, num := range numbers {
		if maxNum < num {
			maxNum = num
		}
	}

	minFuel := math.MaxInt
	var minPosition int

	minFuel2 := math.MaxInt
	var minPosition2 int

	for i := 0; i < maxNum; i++ {
		var fuel int
		var fuel2 int

		for _, pos := range numbers {
			steps := int(math.Abs(float64(i - pos)))

			fuel += steps
			for j := 0; j < steps; j++ {
				fuel2 += j + 1
			}
		}

		if fuel < minFuel {
			minFuel = fuel
			minPosition = i
		}

		if fuel2 < minFuel2 {
			minFuel2 = fuel2
			minPosition2 = i
		}
	}

	fmt.Println("Min fuel required part 1:", minFuel, "at", minPosition)
	fmt.Println("Min fuel required part 2:", minFuel2, "at", minPosition2)
}
