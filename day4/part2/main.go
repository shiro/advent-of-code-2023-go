package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cardCount := 0
	copyCountMap := map[int]int{}

	for scanner.Scan() {
		raw := scanner.Text()
		id, _ := strconv.Atoi(strings.Trim(strings.Split(raw, ": ")[0][len("Card "):], " "))
		fmt.Println("id", id)
		raw = strings.Split(raw, ": ")[1]

		cardCount += 1 + copyCountMap[id]

		parts := strings.Split(raw, " | ")
		winningNumbers := []int{}
		ourNumbers := []int{}

		for _, x := range strings.Split(parts[0], " ") {
			if res, err := strconv.Atoi(x); err == nil {
				winningNumbers = append(winningNumbers, res)
			}
		}
		for _, x := range strings.Split(parts[1], " ") {
			if res, err := strconv.Atoi(x); err == nil {
				ourNumbers = append(ourNumbers, res)
			}
		}

		matches := 0
		for _, number := range ourNumbers {
			if slices.Contains(winningNumbers, number) {
				matches++
				copyCountMap[id+matches] += 1 + copyCountMap[id]
			}
		}
	}

	fmt.Println(cardCount)
}
