package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type mapping struct {
	start        int
	end          int
	mappingDelta int
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	distances := strings.Fields(scanner.Text())[1:]
	result := 1

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		min_distance, _ := strconv.Atoi(distances[i])
		alternatives := 0

		for t := 0; t < time; t++ {
			if (t * (time - t)) > min_distance {
				alternatives++
			}
		}
		fmt.Println("a", alternatives)
		result *= alternatives
	}

	fmt.Println(result)
}
