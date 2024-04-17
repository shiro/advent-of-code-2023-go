package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	for scanner.Scan() {
		raw := scanner.Text()
		parts := strings.Split(raw, ": ")
		parts = strings.Split(parts[1], "; ")

		minSet := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, p := range parts {
			parts := strings.Split(p, ", ")
			for _, p := range parts {
				parts := strings.Split(p, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]

				minSet[color] = max(minSet[color], count)
			}
		}

		power := 1
		for k := range minSet {
			power *= minSet[k]
		}

		sum += power
	}
	fmt.Println(sum)
}
