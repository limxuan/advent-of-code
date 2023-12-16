package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	file, err := os.Open("2023/day2/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxMarbles := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sumOfValidGameIDS := 0

	currentGameID := 1
	for scanner.Scan() {
		isValid := true
		line := scanner.Text()
		excludeGameID := strings.Split(line, ":")[1]
		gameSets := strings.Split(excludeGameID, ";")
	SetsLoop:
		for _, set := range gameSets {
			marbles := strings.Split(set, ",")
			for _, marble := range marbles {
				tMarble := strings.TrimSpace(marble)
				info := strings.Split(tMarble, " ")
				sQuantity, color := info[0], info[1]
				quantity, err := strconv.Atoi(sQuantity)
				if err != nil {
					fmt.Println("is not a number")
				}

				if maxMarbles[color] < quantity {
					isValid = false
					break SetsLoop
				}
			}
		}
		if isValid {
			sumOfValidGameIDS += currentGameID
		}
		currentGameID++
	}

	fmt.Printf("Sum of valid game ids: %d", sumOfValidGameIDS)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
