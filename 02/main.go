package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/fredrikln/advent-of-code-2021/utils"
)

func main() {
	input := ParseFileToStringSlice("input.txt")

	fmt.Println("Part 1:")
	Part1(input)

	fmt.Println("Part 2:")
	Part2(input)
}

func Part1(commands []string) {
	x, y := FollowCommands(commands)

	fmt.Println(x * y)
}

func Part2(commands []string) {
	x, y := FollowCommandsWithAim(commands)

	fmt.Println(x * y)
}

func ParseCommandString(commandString string) (command string, distance int) {
	parts := strings.Split(commandString, " ")

	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("Unable to parse command")
	}

	return parts[0], distance
}

func FollowCommands(commands []string) (x, y int) {
	for _, commandString := range commands {
		command, distance := ParseCommandString(commandString)

		if command == "forward" {
			x += distance
		} else if command == "up" {
			y -= distance
		} else if command == "down" {
			y += distance
		}
	}

	return x, y
}

func FollowCommandsWithAim(commands []string) (x, y int) {
	var aim int

	for _, commandString := range commands {
		command, distance := ParseCommandString(commandString)

		if command == "forward" {
			x += distance
			y += distance * aim
		} else if command == "up" {
			aim -= distance
		} else if command == "down" {
			aim += distance
		}
	}

	return x, y
}
