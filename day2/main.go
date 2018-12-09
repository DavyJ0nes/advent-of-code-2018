package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := getInput("./input.txt")
	if err != nil {
		panic(err)
	}

	output := GenerateChecksum(input)
	fmt.Println(output)

}

// GenerateChecksum iterates through
func GenerateChecksum(input []string) int {
	var left, right int

	for _, s := range input {
		left += countCharsFast(s, 2)
		right += countCharsFast(s, 3)
	}

	return left * right
}

// countCharsFast iterates through a string and returns 1 if the string contains
// the desired number of duplicates
func countCharsFast(s string, want int) int {
	for _, c := range s {
		if count := strings.Count(s, string(c)); count == want {
			return 1
		}
	}

	return 0
}

// countCharsSlow iterates through a string and returns 1 if the string contains
// the desired number of duplicates
func countCharsSlow(s string, want int) int {
	chars := strings.Split(s, "")

	for _, c := range chars {
		if count := strings.Count(s, c); count == want {
			return 1
		}
	}

	return 0
}

// getInput splits the file into lines
func getInput(filename string) ([]string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input := strings.Split(string(body), "\n")

	return input, nil
}
