package main

import (
	"fmt"
	"os"

	"github.com/squarepear/aoc/2025/day1"
	"github.com/squarepear/aoc/2025/day2"
	"github.com/squarepear/aoc/2025/day3"
	"github.com/squarepear/aoc/2025/day4"
)

var Days = map[int]func() error{
	1: day1.Execute,
	2: day2.Execute,
	3: day3.Execute,
	4: day4.Execute,
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the day number to execute")
		return
	}

	day := 0
	_, err := fmt.Sscanf(args[1], "%d", &day)
	if err != nil {
		fmt.Println("Invalid day number")
		return
	}

	if execute, exists := Days[day]; exists {
		fmt.Println("Executing Day", day)
		err := execute()
		if err != nil {
			fmt.Printf("Error executing day %d: %v\n", day, err)
		}
	} else {
		fmt.Println("Day not implemented")
	}
}
