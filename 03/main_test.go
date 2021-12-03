package main

import "testing"

func TestCalculateGammaAndEpsilon(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	gamma, epsilon := CalculateGammaAndEpsilon(input)

	if gamma != 22 {
		t.Errorf("Gamma not corrent, got %d, want %d", gamma, 22)
	}

	if epsilon != 9 {
		t.Errorf("Epsilon not corrent, got %d, want %d", epsilon, 9)
	}
}

func TestCalculateO2AndCO2(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	o2, co2 := CalculateO2AndCO2(input)

	if o2 != 23 {
		t.Errorf("O2 not corrent, got %d, want %d", o2, 23)
	}

	if co2 != 10 {
		t.Errorf("CO2 not corrent, got %d, want %d", co2, 10)
	}
}

func TestFilterByCommonality(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	result := FilterByMostCommon(input, 0)

	if len(result) != 7 {
		t.Errorf("Got wrong amount of lines, got %d, want %d", len(result), 7)
	}

	result2 := FilterByMostCommon(result, 1)

	if len(result2) != 4 {
		t.Errorf("Got wrong amount of lines, got %d, want %d", len(result), 4)
	}
}

func TestFilterByLeastCommon(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	result := FilterByLeastCommon(input, 0)

	if len(result) != 5 {
		t.Errorf("Got wrong amount of lines, got %d, want %d", len(result), 5)
	}

	result2 := FilterByLeastCommon(result, 1)

	if len(result2) != 2 {
		t.Errorf("Got wrong amount of lines, got %d, want %d", len(result), 2)
	}
}
