package day4

import (
	"bufio"
	"fmt"

	"github.com/squarepear/aoc/2025/libary"
)

const paperChar = '@'

type Grid []Row
type Row []int

func Execute() error {
	// Load input file
	input, err := libary.LoadInput(4)
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

func ParseInput(reader *bufio.Reader) (Grid, error) {
	scanner := bufio.NewScanner(reader)
	var inputData Grid

	for scanner.Scan() {
		line := scanner.Text()
		row, err := ParseRow(line)
		if err != nil {
			return nil, err
		}

		inputData = append(inputData, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputData, nil
}

func ParseRow(input string) (Row, error) {
	var row Row

	for _, ch := range input {
		val := 0
		if ch == paperChar {
			val = 1
		}
		row = append(row, val)
	}

	return row, nil
}

func CalculateNeighbors(grid Grid) (Grid, error) {
	var neighborsGrid Grid

	for y, row := range grid {
		var neighborRow Row
		for x := range row {
			neighbors := 0

			neighbors += grid.getCell(x-1, y-1)
			neighbors += grid.getCell(x-1, y)
			neighbors += grid.getCell(x-1, y+1)
			neighbors += grid.getCell(x, y-1)
			neighbors += grid.getCell(x, y+1)
			neighbors += grid.getCell(x+1, y-1)
			neighbors += grid.getCell(x+1, y)
			neighbors += grid.getCell(x+1, y+1)

			neighborRow = append(neighborRow, neighbors)
		}

		neighborsGrid = append(neighborsGrid, neighborRow)
	}

	return neighborsGrid, nil
}

func (g Grid) getCell(x, y int) int {
	if y < 0 || y >= len(g) {
		return 0
	}

	if x < 0 || x >= len(g[y]) {
		return 0
	}

	return g[y][x]
}

func (g Grid) Copy() Grid {
	var grid Grid

	for _, r := range g {
		grid = append(grid, append(Row{}, r...))
	}

	return grid
}

func SolvePart1(data Grid) int {
	solution := 0

	neighborsGrid, err := CalculateNeighbors(data)
	if err != nil {
		return -1
	}

	for y, row := range neighborsGrid {
		for x, neighbors := range row {
			if data[y][x] == 0 {
				neighborsGrid[y][x] = 0
				continue
			}

			if neighbors < 4 {
				solution++
			}
		}
	}

	return solution
}

func SolvePart2(data Grid) int {
	solution := 0

	haveRemoved := true
	for haveRemoved {
		haveRemoved = false
		neighborsGrid, err := CalculateNeighbors(data)
		if err != nil {
			return -1
		}

		data = data.Copy()

		for y, row := range neighborsGrid {
			for x, neighbors := range row {
				if data[y][x] == 0 {
					neighborsGrid[y][x] = 0
					continue
				}

				if neighbors < 4 {
					solution++
					data[y][x] = 0
					haveRemoved = true
				}
			}
		}

	}

	return solution
}
