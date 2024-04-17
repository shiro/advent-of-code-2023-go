package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var seeds []int
	mappings := map[struct {
		from string
		to   string
	}][]mapping{}
	var from, to string
	_, _ = from, to

	for scanner.Scan() {
		raw := scanner.Text()
		if len(seeds) == 0 {
			parts := strings.Split(raw[len("seeds: "):], " ")
			for _, part := range parts {
				parsed, _ := strconv.Atoi(part)
				seeds = append(seeds, parsed)
			}
			continue
		}
		if raw == "" {
			continue
		}

		// parse map name
		if strings.Contains(raw, "map:") {
			parts := strings.Split(strings.Split(raw, " ")[0], "-to-")
			from, to = parts[0], parts[1]
			continue
		}
		parts := strings.Split(raw, " ")
		dst, _ := strconv.Atoi(parts[0])
		src, _ := strconv.Atoi(parts[1])
		range_, _ := strconv.Atoi(parts[2])
		// fmt.Println(dst-src, range_)

		key := struct {
			from string
			to   string
		}{from, to}

		mappings[key] = append(mappings[key], mapping{
			start:        src,
			end:          src + range_,
			mappingDelta: dst - src,
		})
	}

	lookup := func(from string, to string, idx int) int {
		key := struct {
			from string
			to   string
		}{from, to}

		for _, mapping := range mappings[key] {
			if mapping.start <= idx && idx <= mapping.end {
				return idx + mapping.mappingDelta
			}
		}
		return idx
	}

	fullChain := []string{
		"seed",
		"soil",
		"fertilizer",
		"water",
		"light",
		"temperature",
		"humidity",
		"location",
	}

	lowestLocation := math.MaxInt64
	for _, seed := range seeds {
		idx := seed
		chain := fullChain

		for ok := true; ok; ok = len(chain) > 1 {
			idx = lookup(chain[0], chain[1], idx)
			chain = chain[1:]
		}

		lowestLocation = min(lowestLocation, idx)
	}
	fmt.Println(lowestLocation)
}
