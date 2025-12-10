package day4_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/squarepear/aoc/2025/day4"
)

func TestParseInput(t *testing.T) {
	input := ".@.@\n@.@.\n.@.@\n@.@.\n"
	expected := day4.Grid{
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 0, 1, 0},
	}

	reader := bufio.NewReader(strings.NewReader(input))
	result, err := day4.ParseInput(reader)
	if err != nil {
		t.Fatalf("ParseInput returned error: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("ParseInput returned %d entries; want %d", len(result), len(expected))
	}

	for i, r := range result {
		if len(r) != len(expected[i]) {
			t.Fatalf("Entry %d has length %d; want %d", i, len(r), len(expected[i]))
		}
		for j, v := range r {
			if v != expected[i][j] {
				t.Errorf("Entry %d, value %d = %d; want %d", i, j, v, expected[i][j])
			}
		}
	}
}

func TestParseRow(t *testing.T) {
	tests := []struct {
		input    string
		expected day4.Row
	}{
		{".@.@", day4.Row{0, 1, 0, 1}},
		{"@@..", day4.Row{1, 1, 0, 0}},
		{".@@.", day4.Row{0, 1, 1, 0}},
	}

	for _, test := range tests {
		result, err := day4.ParseRow(test.input)
		if err != nil {
			t.Errorf("ParseRow(%q) returned error: %v", test.input, err)
			continue
		}
		if len(result) != len(test.expected) {
			t.Errorf("ParseRow(%q) returned length %d; want %d", test.input, len(result), len(test.expected))
			continue
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("ParseRow(%q)[%d] = %d; want %d", test.input, i, v, test.expected[i])
			}
		}
	}
}

func TestCalculateNeighbors(t *testing.T) {
	input := day4.Grid{
		{0, 1, 0, 1},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 0, 1, 0},
	}
	expected := day4.Grid{
		{2, 2, 3, 1},
		{2, 4, 4, 3},
		{3, 4, 4, 2},
		{1, 3, 2, 2},
	}

	result, err := day4.CalculateNeighbors(input)
	if err != nil {
		t.Fatalf("CalculateNeighbors returned error: %v", err)
	}

	for i, r := range result {
		if len(r) != len(expected[i]) {
			t.Fatalf("Entry %d has length %d; want %d", i, len(r), len(expected[i]))
		}
		for j, v := range r {
			if v != expected[i][j] {
				t.Errorf("Entry %d, value %d = %d; want %d", i, j, v, expected[i][j])
			}
		}
	}
}

func TestSolvePart1(t *testing.T) {
	input := day4.Grid{
		{0, 0, 1, 1, 0, 1, 1, 1, 1, 0},
		{1, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	}
	expected := 13

	result := day4.SolvePart1(input)
	if result != expected {
		t.Errorf("SolvePart1() = %d; want %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	input := day4.Grid{
		{0, 0, 1, 1, 0, 1, 1, 1, 1, 0},
		{1, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 1, 0, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	}
	expected := 43

	result := day4.SolvePart2(input)
	if result != expected {
		t.Errorf("SolvePart2() = %d; want %d", result, expected)
	}
}
