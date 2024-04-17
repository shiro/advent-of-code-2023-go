package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	grid := map[Point]*int{}

	symbolPositions := []Point{}
	y := 0
	for scanner.Scan() {
		raw := scanner.Text()
		buf := ""

		for x := 0; x < len(raw); x++ {
			// collect all sequential numbers into a buffer
			for xOffset := x; xOffset < len(raw); xOffset++ {
				tmp := raw[x : xOffset+1]
				if _, err := strconv.Atoi(tmp); err != nil {
					break
				}
				buf = tmp
			}
			// parse the buffer into a number, update the grid and advance
			if buf != "" {
				number, _ := strconv.Atoi(buf)
				for xOffset := 0; xOffset < len(buf); xOffset++ {
					grid[Point{x + xOffset, y}] = &number
					// fmt.Println(x+xOffset, y, number)
				}
				x += len(buf) - 1
				buf = ""
				continue
			}

			// parse symbol
			ch := string(raw[x])
			if _, err := strconv.Atoi(ch); err != nil && ch != "." {
				symbolPositions = append(symbolPositions, Point{x, y})
			}
		}
		y++
	}

	visitedNumbers := []*int{}
	sum := 0
	for _, pos := range symbolPositions {
		for x := pos.x - 1; x <= pos.x+1; x++ {
			for y := pos.y - 1; y <= pos.y+1; y++ {
				pos := Point{x, y}
				if grid[pos] == nil || slices.Contains(visitedNumbers, grid[pos]) {
					continue
				}
				visitedNumbers = append(visitedNumbers, grid[pos])
				sum += *grid[pos]
			}
		}
	}
	fmt.Println(sum)
}
