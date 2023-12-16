package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("2023/day2/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPower := 0
	for scanner.Scan() {
		line := scanner.Text()
		excludeGameID := strings.Split(line, ":")[1]
		gameSets := strings.Split(excludeGameID, ";")
		minMarbles := map[string]int{}

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

				if minMarbles[color] < quantity {
					minMarbles[color] = quantity
				}
			}
		}
		power := minMarbles["red"] * minMarbles["green"] * minMarbles["blue"]
		fmt.Println("power", power)
		totalPower += power
	}

	fmt.Printf("total power:%d\n", totalPower)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
