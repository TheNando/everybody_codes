package quests

import (
	"errors"
	"strconv"

	"everybody-codes/internal/util"
)

func GetName(names []string, instructions []string) (string, error) {
	nameIndex := 0
	maxIndex := len(names) - 1

	for _, instruction := range instructions {
		dir := string(instruction[0])
		steps, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return "", errors.New("failed to parse instruction")
		}

		switch dir {
		case "L":
			nameIndex -= steps
		case "R":
			nameIndex += steps
		}

		if nameIndex < 0 {
			nameIndex = 0
		} else if nameIndex > maxIndex {
			nameIndex = maxIndex
		}

	}

	name := names[nameIndex]

	return name, nil
}

func GetCycleName(names []string, instructions []string) (string, error) {
	nameCycle := util.NewCycle(names)

	for _, instruction := range instructions {
		dir := string(instruction[0])
		steps, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return "", errors.New("failed to parse instruction")
		}

		switch dir {
		case "L":
			nameCycle.Left(steps)
		case "R":
			nameCycle.Right(steps)
		}
	}

	return nameCycle.Item(), nil
}

func SwapCycleName(names []string, instructions []string) (string, error) {
	nameCycle := util.NewCycle(names)

	for _, instruction := range instructions {
		dir := string(instruction[0])
		steps, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return "", errors.New("failed to parse instruction")
		}

		nameCycle.Swap(dir, steps)
	}

	return nameCycle.Item(), nil
}
