package main

import (
	"fmt"

	. "github.com/fredrikln/advent-of-code-2021/utils"
)

// type lanternfish struct {
// 	timer int
// }

func main() {
	numbers := ParseFirstRowToIntSlice("input.txt")

	// Initial solution for step 1 that does not work for step 2
	// school := make([]lanternfish, 0, len(numbers))

	// for _, number := range numbers {
	// 	school = append(school, lanternfish{number})
	// }

	// fmt.Println(school)

	// for i := 0; i < 80; i++ {
	// 	var newFish int
	// 	for i := range school {
	// 		fish := &school[i]

	// 		fish.timer--

	// 		if fish.timer == -1 {
	// 			fish.timer = 6
	// 			newFish++
	// 		}
	// 	}

	// 	for j := 0; j < newFish; j++ {
	// 		school = append(school, lanternfish{8})
	// 	}

	// 	// fmt.Println(school)
	// }

	// fmt.Println("Num fish part 1:", len(school))

	// Better solution that works for step 2
	days := make([]int, 256+9)

	// fmt.Println(numbers)
	for _, number := range numbers {
		days[number]++

		for i := number + 7; i < 256; i += 7 {
			days[i]++
		}
	}

	steps := 256

	for i := 0; i < steps; i++ {
		days[i+9] += days[i]

		for j := i + 9 + 7; j < 256; j += 7 {
			days[j] += days[i]
		}
	}

	// fmt.Println(days[:steps])

	sum := len(numbers)
	for i := 0; i < steps; i++ {
		if i == 80 {
			fmt.Println("Sum fish part 1:", sum)
		}
		sum += days[i]
	}

	fmt.Println("Sum fish part 2:", sum)
}
