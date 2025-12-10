package day5

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/squarepear/aoc/2025/libary"
)

type Data struct {
	IDRanges    []IDRange
	Ingredients []Ingredient
}

type IDRange struct {
	Start, End int
}
type Ingredient struct {
	ID int
}

func Execute() error {
	// Load input file
	input, err := libary.LoadInput(5)
	if err != nil {
		return err
	}

	// Parse input data
	data, err := ParseInput(input)
	if err != nil {
		return err
	}

	// Solve part 1
	part1Result := SolvePart1(data)
	fmt.Printf("Part 1 Result: %d\n", part1Result)

	// Solve part 2
	part2Result := SolvePart2(data)
	fmt.Printf("Part 2 Result: %d\n", part2Result)

	return nil
}

func ParseInput(reader *bufio.Reader) (Data, error) {
	scanner := bufio.NewScanner(reader)
	var inputData Data

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}

		idRange, err := ParseIDRange(line)
		if err != nil {
			return Data{}, err
		}

		inputData.IDRanges = append(inputData.IDRanges, idRange)
	}

	for scanner.Scan() {
		line := scanner.Text()
		ingredient, err := ParseIngredient(line)
		if err != nil {
			return Data{}, err
		}

		inputData.Ingredients = append(inputData.Ingredients, ingredient)
	}

	if err := scanner.Err(); err != nil {
		return Data{}, err
	}

	return inputData, nil
}

// Range format: "X-Y" where X and Y are integers
func ParseIDRange(input string) (IDRange, error) {
	var idRange IDRange

	_, err := fmt.Sscanf(input, "%d-%d", &idRange.Start, &idRange.End)
	if err != nil {
		return IDRange{}, err
	}

	if idRange.End < idRange.Start {
		return IDRange{}, fmt.Errorf("invalid range: end %d is less than start %d", idRange.End, idRange.Start)
	}

	return idRange, nil
}

func ParseIngredient(input string) (Ingredient, error) {
	var ingredient Ingredient

	var err error
	ingredient.ID, err = strconv.Atoi(input)
	if err != nil {
		return Ingredient{}, err
	}

	return ingredient, nil
}

func MergeOverlaps(rangesa []IDRange) []IDRange {
	merged := append([]IDRange{}, rangesa...)

	for i := 0; i < len(merged); i++ {
		for j := i + 1; j < len(merged); j++ {
			overlaps := merged[i].Start <= merged[j].End+1 && merged[j].Start <= merged[i].End+1

			if overlaps {
				if merged[j].Start < merged[i].Start {
					merged[i].Start = merged[j].Start
				}
				if merged[j].End > merged[i].End {
					merged[i].End = merged[j].End
				}

				merged = append(merged[:j], merged[j+1:]...)
				j = i
			}
		}
	}

	return merged
}

func SolvePart1(data Data) int {
	solution := 0

	for _, ingredient := range data.Ingredients {
		for _, idRange := range data.IDRanges {
			if ingredient.ID >= idRange.Start && ingredient.ID <= idRange.End {
				solution++
				break
			}
		}
	}

	return solution
}

func SolvePart2(data Data) int {
	solution := 0

	ranges := MergeOverlaps(data.IDRanges)

	for _, idRange := range ranges {
		solution += idRange.End - idRange.Start + 1
	}

	return solution
}
