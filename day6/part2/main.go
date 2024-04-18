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
	time, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))
	scanner.Scan()
	min_distance, _ := strconv.Atoi(strings.Join(strings.Fields(scanner.Text())[1:], ""))
	result := 1
	fmt.Println(time, min_distance)

	alternatives := 0

	for t := 0; t < time; t++ {
		if (t * (time - t)) > min_distance {
			alternatives++
		}
	}
	result *= alternatives

	fmt.Println(result)
}
