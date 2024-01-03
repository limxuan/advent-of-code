package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func PartOne() {
	file, err := os.Open("2023/day4/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}
	totalPoints := float64(0)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
		excludeCardID := strings.Split(line, ":")[1]
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
			totalPoints += math.Pow(2, (float64(correctNumbers) - 1))
		}
	}
	fmt.Printf("Total points: %f", totalPoints)
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
