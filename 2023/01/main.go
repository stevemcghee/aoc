package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new scanner.
	scanner := bufio.NewScanner(f)

	// Loop over the lines.
	for scanner.Scan() {
		// Get the line.
		line := scanner.Text()

		var firstNumber = 0
		var lastNumber = 0

		// Split the line into chars.
		chars := strings.Split(line, "")
		for _, c := range chars {
			if isNumber(byte(c)) == false {
				continue
			}
			// Find the first and last number.
			firstNumber, err := strconv.Atoi(chars[0])
			if err != nil {
				log.Fatal(err)
			}
			lastNumber, err := strconv.Atoi(chars[len(chars)-1])
			if err != nil {
				log.Fatal(err)
			}
		}

		// Print the first and last number.
		fmt.Println(firstNumber, lastNumber)
	}
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}