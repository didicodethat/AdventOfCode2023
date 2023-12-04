package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var strToNumber = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2, "3": 3,
		"4": 4, "5": 5,
		"6": 6, "7": 7,
		"8": 8, "9": 9,
	}
	var finalSum uint64 = 0
	readFile, err := os.Open("./day1.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		const unitializedValue = -1
		number := []int{0, 0}
		var hits []int
		for i := 0; i < len(line); i++ {
			currentSubstring := line[i:]
			for key, value := range strToNumber {
				if key == currentSubstring[0:min(len(key), len(currentSubstring))] {
					hits = append(hits, value)
				}
			}
		}
		number[0] = hits[0]
		number[1] = hits[len(hits)-1]
		finalSum += uint64(number[0]*10 + number[1])
	}
	fmt.Println(finalSum)
	readFile.Close()
}
