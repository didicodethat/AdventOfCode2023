package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numericalCharacters = "0123456789"

func main() {
	var finalSum uint64 = 0
	readFile, err := os.Open("./day1.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		number := []string{"", ""}
		chars := strings.Split(line, "")
		for _, currentChar := range chars {
			if strings.Contains(numericalCharacters, currentChar) {
				if number[0] == "" {
					number[0] = currentChar
				}
				number[1] = currentChar
			}
		}
		currentNumber, convertionError := strconv.Atoi(strings.Join(number, ""))
		if convertionError != nil {
			panic(convertionError)
		}
		finalSum += uint64(currentNumber)
	}
	fmt.Println(finalSum)
	readFile.Close()
}
