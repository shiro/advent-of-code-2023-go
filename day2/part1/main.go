package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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
		id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
		parts = strings.Split(parts[1], "; ")
		possible := true

	out:
		for _, p := range parts {
			parts := strings.Split(p, ", ")
			for _, p := range parts {
				parts := strings.Split(p, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]
				if limits[color] < count {
					possible = false
					break out
				}
			}
		}
		if possible {
			sum += id
		}
	}
	fmt.Println(sum)
}
