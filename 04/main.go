package main

import (
	"fmt"
	. "github.com/fredrikln/advent-of-code-2021/utils"
	"strconv"
	"strings"
)

func main() {
    rows := ParseFileToStringSlice("input.txt")

	var calls string
	var board *BingoBoard
	var boards []*BingoBoard

	for index, row := range rows {
		if index == 0 {
			calls = row
			continue
		}

		if row == "" {
			board = &BingoBoard{}
			boards = append(boards, board)
			continue
		}

		numbers := strings.Fields(row)

		for _, number := range numbers {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			bingoNumber := BingoNumber{n, false}
			board.Numbers = append(board.Numbers, bingoNumber)
		}
	}

	fmt.Println("Calls:", calls)
	callsLoop:
		for ni, number := range strings.Split(calls, ",") {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			for i, board := range boards {
				board.Pull(n)

				if board.HasBingo() {
					fmt.Println("Board", i + 1, "has bingo on number", ni + 1, "of", len(strings.Split(calls, ",")))
					fmt.Println(board)

					var total int
					for _, boardNumber := range board.Numbers {
						if !boardNumber.Pulled {
							total += boardNumber.Number
						}
					}

					fmt.Println("Score Part 1:", total * n)

					break callsLoop
				}
			}
		}

	wonBoardsMap := make(map[int]bool)

	fmt.Println()

	callsLoop2:
		for ni, number := range strings.Split(calls, ",") {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			for i, board := range boards {
				board.Pull(n)

				if board.HasBingo() {
					exists := wonBoardsMap[i]

					if !exists {
						wonBoardsMap[i] = true
					}

					if len(wonBoardsMap) == len(boards) {
						fmt.Println("Board", i + 1, "has bingo on number", ni + 1, "of", len(strings.Split(calls, ",")))
						fmt.Println(board)


						var total int
						for _, boardNumber := range board.Numbers {
							if !boardNumber.Pulled {
								total += boardNumber.Number
							}
						}

						fmt.Println("Score Part 2:", total * n)

						break callsLoop2
					}
				}
			}
		}
}

func Part1() {

}

func Part2() {

}

type BingoNumber struct {
    Number int
    Pulled bool
}

type BingoBoard struct {
    Numbers []BingoNumber
}

func (bb *BingoBoard) HasBingo() bool {
	for i := 0; i < 5; i++ {
		// rows
		if bb.Numbers[0+i*5].Pulled && bb.Numbers[1+i*5].Pulled && bb.Numbers[2+i*5].Pulled && bb.Numbers[3+i*5].Pulled && bb.Numbers[4+i*5].Pulled {
			return true
		}

		// cols
		if bb.Numbers[0+i].Pulled && bb.Numbers[5+i].Pulled && bb.Numbers[10+i].Pulled && bb.Numbers[15+i].Pulled && bb.Numbers[20+i].Pulled {
			return true
		}
	}

	return false
}

func (bb *BingoBoard) String() string {
	out := ""

	for i := 0; i < len(bb.Numbers); i++ {
		if i % 5 == 0 {
			fmt.Println()
		}

		if bb.Numbers[i].Pulled {
			fmt.Print("*")
		}

		fmt.Printf("%v ", bb.Numbers[i].Number)
	}

	return out
}

func (bb *BingoBoard) Pull(number int) {
	for i := 0; i < len(bb.Numbers); i++ {
		if bb.Numbers[i].Number == number {
			bb.Numbers[i].Pulled = true
		}
	}
}
