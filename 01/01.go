package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/smcghee/src/aoc/01/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
		input = append(input, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	numLen := len(input)

	for i := 0; i < numLen; i++ {
		for j := i; j < numLen; j++ {
			if input[i]+input[j] == 2020 {
				fmt.Println(input[i], "and", input[j])
				fmt.Println(" ==> ", input[i]*input[j])
			}
		}
	}
}
