package day5_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/squarepear/aoc/2025/day5"
)

func TestParseInput(t *testing.T) {
	input := "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32\n"
	expected := day5.Data{
		IDRanges: []day5.IDRange{
			{Start: 3, End: 5},
			{Start: 10, End: 14},
			{Start: 16, End: 20},
			{Start: 12, End: 18},
		},
		Ingredients: []day5.Ingredient{
			{ID: 1},
			{ID: 5},
			{ID: 8},
			{ID: 11},
			{ID: 17},
			{ID: 32},
		},
	}

	reader := bufio.NewReader(strings.NewReader(input))
	result, err := day5.ParseInput(reader)
	if err != nil {
		t.Fatalf("ParseInput returned error: %v", err)
	}

	if len(result.IDRanges) != len(expected.IDRanges) {
		t.Fatalf("ParseInput returned %d IDRanges; want %d", len(result.IDRanges), len(expected.IDRanges))
	}
	for i, r := range result.IDRanges {
		if r != expected.IDRanges[i] {
			t.Errorf("IDRange %d = %v; want %v", i, r, expected.IDRanges[i])
		}
	}

	if len(result.Ingredients) != len(expected.Ingredients) {
		t.Fatalf("ParseInput returned %d Ingredients; want %d", len(result.Ingredients), len(expected.Ingredients))
	}
	for i, ing := range result.Ingredients {
		if ing != expected.Ingredients[i] {
			t.Errorf("Ingredient %d = %v; want %v", i, ing, expected.Ingredients[i])
		}
	}
}

func TestParseRange(t *testing.T) {
	tests := []struct {
		input    string
		expected day5.IDRange
	}{
		{"1-3", day5.IDRange{Start: 1, End: 3}},
		{"5-7", day5.IDRange{Start: 5, End: 7}},
		{"10-15", day5.IDRange{Start: 10, End: 15}},
	}

	for _, test := range tests {
		result, err := day5.ParseIDRange(test.input)
		if err != nil {
			t.Errorf("ParseRange(%q) returned error: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("ParseRange(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestParseRange_InvalidInput(t *testing.T) {
	var tests = []string{
		"3-1",
		"abc-def",
		"5-",
		"-10",
	}

	for _, input := range tests {
		_, err := day5.ParseIDRange(input)
		if err == nil {
			t.Errorf("ParseRange(%q) expected error but got nil", input)
		}
	}
}

func TestParseIngredient(t *testing.T) {
	tests := []struct {
		input    string
		expected day5.Ingredient
	}{
		{"1", day5.Ingredient{ID: 1}},
		{"42", day5.Ingredient{ID: 42}},
		{"100", day5.Ingredient{ID: 100}},
	}

	for _, test := range tests {
		result, err := day5.ParseIngredient(test.input)
		if err != nil {
			t.Errorf("ParseIngredient(%q) returned error: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("ParseIngredient(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestMergeOverlaps(t *testing.T) {
	tt := []struct {
		name     string
		input    []day5.IDRange
		expected []day5.IDRange
	}{
		{
			"overlapping ranges",
			[]day5.IDRange{
				{Start: 1, End: 3},
				{Start: 2, End: 5},
			},
			[]day5.IDRange{
				{Start: 1, End: 5},
			},
		},
		{
			"adjacent ranges",
			[]day5.IDRange{
				{Start: 1, End: 3},
				{Start: 5, End: 8},
			},
			[]day5.IDRange{
				{Start: 1, End: 3},
				{Start: 5, End: 8},
			},
		},
		{
			"directly adjacent ranges",
			[]day5.IDRange{
				{Start: 1, End: 3},
				{Start: 4, End: 6},
			},
			[]day5.IDRange{
				{Start: 1, End: 6},
			},
		},
		{
			"multiple overlapping ranges",
			[]day5.IDRange{
				{Start: 1, End: 4},
				{Start: 3, End: 6},
				{Start: 5, End: 8},
			},
			[]day5.IDRange{
				{Start: 1, End: 8},
			},
		},
		{
			"mixed ranges",
			[]day5.IDRange{
				{Start: 1, End: 2},
				{Start: 4, End: 5},
				{Start: 3, End: 6},
				{Start: 8, End: 10},
				{Start: 9, End: 12},
			},
			[]day5.IDRange{
				{Start: 1, End: 6},
				{Start: 8, End: 12},
			},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			result := day5.MergeOverlaps(test.input)
			if len(result) != len(test.expected) {
				t.Fatalf("MergeOverlaps returned %d ranges; want %d", len(result), len(test.expected))
			}
			for i, r := range result {
				if r != test.expected[i] {
					t.Errorf("Range %d = %v; want %v", i, r, test.expected[i])
				}
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	input := day5.Data{
		IDRanges: []day5.IDRange{
			{Start: 3, End: 5},
			{Start: 10, End: 14},
			{Start: 16, End: 20},
			{Start: 12, End: 18},
		},
		Ingredients: []day5.Ingredient{
			{ID: 1},
			{ID: 5},
			{ID: 8},
			{ID: 11},
			{ID: 17},
			{ID: 32},
		},
	}
	expected := 3

	result := day5.SolvePart1(input)
	if result != expected {
		t.Errorf("SolvePart1() = %d; want %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	input := day5.Data{
		IDRanges: []day5.IDRange{
			{Start: 3, End: 5},
			{Start: 10, End: 14},
			{Start: 16, End: 20},
			{Start: 12, End: 18},
		},
	}
	expected := 14

	result := day5.SolvePart2(input)
	if result != expected {
		t.Errorf("SolvePart2() = %d; want %d", result, expected)
	}
}
