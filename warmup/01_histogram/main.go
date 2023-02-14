package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type runes []rune

func main() {
	lines, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}
	counter := countSymbols(lines)
	printHistogram(counter)
}

func (r runes) Len() int {
	return len(r)
}

func (r runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r runes) Less(i, j int) bool {
	return r[i] < r[j]
}

func readInput() ([]string, error) {
	res := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func countSymbols(lines []string) map[rune]int {
	if len(lines) == 0 {
		return nil
	}

	counter := make(map[rune]int)
	for _, line := range lines {
		for _, sym := range line {
			if sym != ' ' {
				if _, ok := counter[sym]; !ok {
					counter[sym] = 1
				} else {
					counter[sym]++
				}
			}
		}
	}

	return counter
}

func printHistogram(counter map[rune]int) {
	syms := make(runes, 0, len(counter))
	for sym := range counter {
		syms = append(syms, sym)
	}
	sort.Sort(syms)
	hist(syms, counter)
}

func hist(syms runes, counter map[rune]int) {
	maxCount := counter[syms[0]]
	for _, count := range counter {
		if maxCount < count {
			maxCount = count
		}
	}

	histogram := make([]string, 0, maxCount)

	var sb strings.Builder

	for i := 0; i < maxCount; i++ {
		for _, sym := range syms {
			if counter[sym] > 0 {
				counter[sym]--
				sb.WriteRune('#')
			} else {
				sb.WriteRune(' ')
			}
		}
		histogram = append(histogram, sb.String())
		sb.Reset()
	}

	for i := range histogram {
		fmt.Println(histogram[maxCount-i-1])
	}

	for i := range syms {
		fmt.Printf("%s", string(syms[i]))
	}
}
