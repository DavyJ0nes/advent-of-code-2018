package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := GetInput("./input.txt")
	if err != nil {
		panic(err)
	}

	total, err := CalcTotal(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(total)
}

// CalcTotal iterates through the input and sums all values
func CalcTotal(input []string) (int, error) {
	total := 0
	for _, num := range input {
		if num == "" {
			continue
		}

		i, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}

		total += i
	}

	return total, nil
}

// GetInput splits the file into lines
func GetInput(filename string) ([]string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(body), "\n")

	return input, nil
}
