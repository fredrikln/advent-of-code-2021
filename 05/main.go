package main

import (
	"fmt"

	. "github.com/fredrikln/advent-of-code-2021/utils"
)

func main() {
	rows := ParseFileToStringSlice("input.txt")

	grid := make(map[string]int)
	grid2 := make(map[string]int)

	var maxX int = 0
	var maxY int = 0

	for _, row := range rows {
		var x1, y1, x2, y2 int

		fmt.Sscanf(row, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if maxX < x1 {
			maxX = x1
		}

		if maxX < x2 {
			maxX = x2
		}

		if maxY < y1 {
			maxY = y1
		}

		if maxY < y2 {
			maxY = y2
		}

		if x1 == x2 || y1 == y2 {
			if x1 == x2 {
				if y1 > y2 {
					y1, y2 = y2, y1
				}

				for i := y1; i <= y2; i++ {
					grid[fmt.Sprintf("%d,%d", x1, i)]++
					grid2[fmt.Sprintf("%d,%d", x1, i)]++
				}
			}

			if y1 == y2 {
				if x1 > x2 {
					x1, x2 = x2, x1
				}

				for i := x1; i <= x2; i++ {
					grid[fmt.Sprintf("%d,%d", i, y1)]++
					grid2[fmt.Sprintf("%d,%d", i, y1)]++
				}
			}
		} else {
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}

			for y, x := y1, x1; x <= x2; {
				grid2[fmt.Sprintf("%d,%d", x, y)]++

				x++

				if y1 > y2 {
					y--
				} else {
					y++
				}
			}
		}
	}

	// for y := 0; y <= maxY; y++ {
	// 	for x := 0; x <= maxX; x++ {
	// 		fmt.Printf("%3d", grid[fmt.Sprintf("%d,%d", x, y)])
	// 	}
	// 	fmt.Println()
	// }

	var count int
	for _, value := range grid {
		if value > 1 {
			count++
		}
	}

	fmt.Println("Part 1 count:", count)

	// for y := 0; y <= maxY; y++ {
	// 	for x := 0; x <= maxX; x++ {
	// 		fmt.Printf("%3d", grid2[fmt.Sprintf("%d,%d", x, y)])
	// 	}
	// 	fmt.Println()
	// }

	var count2 int
	for _, value := range grid2 {
		if value > 1 {
			count2++
		}
	}

	fmt.Println("Part 2 count:", count2)
}
