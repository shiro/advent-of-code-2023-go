package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	from string
	to   int
}

var numbers = [...]struct {
	string
	int
}{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		raw := scanner.Text()
		var first, last int

	out1:
		for i := 0; i < len(raw); i++ {
			intval := int(raw[i] - '0')
			if 0 <= intval && intval <= 9 {
				first = intval
				break
			}
			for _, v := range numbers {
				if strings.HasPrefix(raw[i:], v.string) {
					first = v.int
					break out1
				}
			}
		}

	out2:
		for i := len(raw) - 1; i >= 0; i-- {
			intval := int(raw[i] - '0')
			if 0 <= intval && intval <= 9 {
				last = intval
				break
			}
			for _, v := range numbers {
				if strings.HasSuffix(raw[:i+1], v.string) {
					last = v.int
					break out2
				}
			}
		}

		res, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		if err != nil {
			log.Fatal(err)
		}

		sum += res
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
