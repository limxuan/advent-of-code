package day4

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func PartTwo() {
	file, err := os.Open("2023/day4/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}
	totalCardIds := 0

	additionalCards := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()
		totalCardIds++
		input = append(input, line)
		cardIDStr, excludeCardID := strings.Fields(strings.Split(line, ":")[0])[1], strings.Split(line, ":")[1]
		cardID, _ := strconv.Atoi(cardIDStr)
		sections := strings.Split(excludeCardID, "|")
		winningNumbersUnparsed, inputNumbersUnparsed := sections[0], sections[1]
		winningNumbers := strings.Fields(winningNumbersUnparsed)
		inputNumbers := strings.Fields(inputNumbersUnparsed)

		correctNumbers := 0

		checked := []string{}
		for _, winningNumber := range winningNumbers {
			for _, inputNumber := range inputNumbers {
				if slices.Contains(checked, inputNumber) {
					continue
				}
				if winningNumber == inputNumber {
					checked = append(checked, winningNumber)
					correctNumbers++
				}
			}
		}

		if correctNumbers != 0 {
			currentQuantity := additionalCards[cardID] + 1
			for i := cardID + 1; i <= cardID+correctNumbers; i++ {
				additionalCards[i] += currentQuantity
			}
		}

	}

	totalScratchcards := 0
	for i := 0; i < totalCardIds; i++ {
		totalScratchcards += additionalCards[i+1] + 1
	}

	fmt.Printf("totalScratchcards: %v\n", totalScratchcards)
}
