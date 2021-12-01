package main

import "testing"

func TestGetNumIncreasing(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"case 1",
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			7,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NumIncreasing(tc.input)

			if got != tc.want {
				t.Errorf("Invalid count, got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestGetNumIncreasingSlidingWindow(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"case 1",
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NumIncreasingSlidingWindow(tc.input)

			if got != tc.want {
				t.Errorf("Invalid count, got %d, want %d", got, tc.want)
			}
		})
	}
}
