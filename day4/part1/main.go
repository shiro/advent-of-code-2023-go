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

	sum := 0
	for scanner.Scan() {
		raw := scanner.Text()
		raw = strings.Split(raw, ": ")[1]

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

		score := 0
		for _, number := range ourNumbers {
			if slices.Contains(winningNumbers, number) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		sum += score
	}

	fmt.Println(sum)
}
