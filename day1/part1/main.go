package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
		var first, last int

		for i := 0; i < len(raw); i++ {
			intval := int(raw[i] - '0')
			if 0 <= intval && intval <= 9 {
				first = intval
				break
			}
		}

		for i := len(raw) - 1; i >= 0; i-- {
			intval := int(raw[i] - '0')
			if 0 <= intval && intval <= 9 {
				last = intval
				break
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
