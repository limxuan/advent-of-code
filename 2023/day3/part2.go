package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getCoordinates(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

var gearMap = map[string][]int{}

func PartTwo() {
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

	// Gear checking
	for rowIdx, row := range input {
		for colIdx := 0; colIdx < len(row); colIdx++ {
			char := row[colIdx]
			isGear := string(char) == "*"
			if isGear {
				gearMap[getCoordinates(colIdx, rowIdx)] = []int{}
			}
		}
	}

	fmt.Printf("%q\n", gearMap)

	for rowIdx, row := range input {
		consecutiveNumber := ""
		connectedTo := ""
		for colIdx := 0; colIdx < len(row); colIdx++ {
			fmt.Println(consecutiveNumber)
			char := row[colIdx]
			if !unicode.IsNumber(rune(char)) || colIdx == len(row)-1 {
				if unicode.IsNumber(rune(char)) {
					if connectedTo == "" {
						connectedTo = isConnectedToGear(rowIdx, colIdx, len(row), len(input), consecutiveNumber)
					}
					consecutiveNumber = fmt.Sprintf("%s%s", consecutiveNumber, string(char))
				}
				if connectedTo != "" {
					iConsecutiveNumber, _ := strconv.Atoi(consecutiveNumber)
					gearMap[connectedTo] = append(gearMap[connectedTo], iConsecutiveNumber)
				}
				consecutiveNumber = ""
				connectedTo = ""
				continue
			} else {
				if connectedTo == "" {
					isConnected := isConnectedToGear(rowIdx, colIdx, len(row), len(input), consecutiveNumber)
					if len(isConnected) > 0 {
						connectedTo = isConnected
					}
				}
				consecutiveNumber = fmt.Sprintf("%s%s", consecutiveNumber, string(char))
			}
		}
	}
	totalSumOfGears := 0
	for _, v := range gearMap {
		if len(v) != 2 {
			continue
		}
		fmt.Printf("%v\n", v)
		gearRatio := v[0] * v[1]
		totalSumOfGears += gearRatio
	}

	fmt.Println(totalSumOfGears)
}

// Returns
//   - connectedTo: returns the coordinate that the number is connected to
func isConnectedToGear(rowIdx, colIdx, maxColLength, maxRowLength int, consecutiveNumber string) (connectedTo string) {
	coordinatesToSearch := [][2]int{}

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
		x, y := coordinate[0], coordinate[1]
		_, exists := gearMap[getCoordinates(x, y)]
		if exists {
			connectedTo = getCoordinates(x, y)
			return
		} else {
			connectedTo = ""
		}
	}
	return
}
