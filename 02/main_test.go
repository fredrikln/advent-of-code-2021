package main

import "testing"

func TestFollowCommands(t *testing.T) {
	cases := []struct {
		name     string
		commands []string
		x        int
		y        int
	}{
		{
			"case 1",
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			15,
			10,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotX, gotY := FollowCommands(tc.commands)

			if gotX != tc.x {
				t.Errorf("Invalid X coordinate, got %d, want %d", gotX, tc.x)
			}
			if gotY != tc.y {
				t.Errorf("Invalid Y coordinate, got %d, want %d", gotY, tc.y)
			}
		})
	}
}

func TestFollowCommandsWithAim(t *testing.T) {
	cases := []struct {
		name     string
		commands []string
		x        int
		y        int
	}{
		{
			"case 1",
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			15,
			60,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotX, gotY := FollowCommandsWithAim(tc.commands)

			if gotX != tc.x {
				t.Errorf("Invalid X coordinate, got %d, want %d", gotX, tc.x)
			}
			if gotY != tc.y {
				t.Errorf("Invalid Y coordinate, got %d, want %d", gotY, tc.y)
			}
		})
	}
}

func TestParseCommandString(t *testing.T) {
	cases := []struct {
		name          string
		commandString string
		command       string
		distance      int
	}{
		{
			"case 1",
			"forward 5",
			"forward",
			5,
		},
		{
			"case 2",
			"down 10",
			"down",
			10,
		},
		{
			"case 3",
			"up 9",
			"up",
			9,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			command, distance := ParseCommandString(tc.commandString)

			if command != tc.command {
				t.Errorf("Invalid command, got %s, want %s", command, tc.command)
			}
			if distance != tc.distance {
				t.Errorf("Invalid distance, got %d, want %d", distance, tc.distance)
			}
		})
	}
}
