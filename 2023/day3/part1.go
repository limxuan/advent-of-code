package day3

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

var symbolMatrix = [][]int{}

func PartOne() {
	file, err := os.Open("2023/day3/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	symbolMatrix = make([][]int, len(input))

	// Symbol checking
	for rowIdx, row := range input {
		for colIdx := 0; colIdx < len(row); colIdx++ {
			char := row[colIdx]
			isSymbol := !unicode.IsNumber(rune(char)) && string(char) != "."
			if isSymbol {
				symbolMatrix[rowIdx] = append(symbolMatrix[rowIdx], colIdx)
			}
		}
	}

	// Number checking
	sumOfValidNumbers := 0
	for rowIdx, row := range input {
		consecutiveNumber := ""
		isValid := false
		for colIdx := 0; colIdx < len(row); colIdx++ {
			char := row[colIdx]
			if !unicode.IsNumber(rune(char)) || colIdx == len(row)-1 {
				if unicode.IsNumber(rune(char)) {
					consecutiveNumber = fmt.Sprintf("%s%s", consecutiveNumber, string(char))
				}
				if isValid {
					iConsecutiveNumber, _ := strconv.Atoi(consecutiveNumber)
					fmt.Println(consecutiveNumber)
					sumOfValidNumbers += iConsecutiveNumber
				}
				consecutiveNumber = ""
				isValid = false
				continue
			} else {
				if !isValid {
					isAdjacent := isAdjacentToSymbol(rowIdx, colIdx, len(row), len(input), consecutiveNumber)
					if isAdjacent {
						isValid = true
					}
				}
			}

		}
	}

	fmt.Printf("Sum of valid numbers: %d", sumOfValidNumbers)
}

func isAdjacentToSymbol(rowIdx, colIdx, maxColLength, maxRowLength int, consecutiveNumber string) (isAdjacent bool) {
	coordinatesToSearch := [][2]int{}
	isAdjacent = false

	appendValidCoordinates := func(rowToSearch int) {
		coordinatesToSearch = append(coordinatesToSearch, [2]int{colIdx, rowToSearch})
		if colIdx != 0 {
			coordinatesToSearch = append(coordinatesToSearch, [2]int{colIdx - 1, rowToSearch})
		}
		if colIdx != maxColLength {
			coordinatesToSearch = append(coordinatesToSearch, [2]int{colIdx + 1, rowToSearch})
		}
	}

	// Search above
	if rowIdx != 0 {
		appendValidCoordinates(rowIdx - 1)
	}

	// Search current row
	appendValidCoordinates(rowIdx)

	// Search below
	if rowIdx != maxRowLength-1 {
		appendValidCoordinates(rowIdx + 1)
	}

	for _, coordinate := range coordinatesToSearch {
		hasSymbol := slices.Contains(symbolMatrix[coordinate[1]], coordinate[0])
		if hasSymbol {
			isAdjacent = true
			break
		}
	}

	return
}
