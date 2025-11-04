package main

import (
	"fmt"
	"os"
	"strings"

	"everybody-codes/internal/quests"
	"everybody-codes/internal/util"
)

func main() {
	// Part 1
	lines, err := util.ReadLines("./input/quest1-part1.txt")
	if err != nil {
		fmt.Printf("Failed to read input: %v\n", err)
		os.Exit(1)
	}

	names := strings.Split(lines[0], ",")
	instructions := strings.Split(lines[1], ",")

	name, err := quests.GetName(names, instructions)
	if err != nil {
		fmt.Printf("Failed to get instructed name: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Your name is " + name)

	// Part 2
	lines, err = util.ReadLines("./input/quest1-part2.txt")
	if err != nil {
		fmt.Printf("Failed to read input: %v\n", err)
		os.Exit(1)
	}

	names = strings.Split(lines[0], ",")
	instructions = strings.Split(lines[1], ",")

	name, err = quests.GetCycleName(names, instructions)
	if err != nil {
		fmt.Printf("Failed to get instructed name: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Your first parent's name is " + name)

	// Part 3
	lines, err = util.ReadLines("./input/quest1-part3.txt")
	if err != nil {
		fmt.Printf("Failed to read input: %v\n", err)
		os.Exit(1)
	}

	names = strings.Split(lines[0], ",")
	instructions = strings.Split(lines[1], ",")

	name, err = quests.SwapCycleName(names, instructions)
	if err != nil {
		fmt.Printf("Failed to get instructed name: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Your second parent's name is " + name)
}
