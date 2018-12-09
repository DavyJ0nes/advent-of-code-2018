package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	sumInputOne()
	sumInputTwo()
}

func sumInputOne() {
	input, err := GetInput("./input.txt")
	if err != nil {
		panic(err)
	}

	total, err := CalcTotal(input...)
	if err != nil {
		panic(err)
	}

	fmt.Println("first iteration total: ", total)

	result := DuplicateFreq(input)
	fmt.Println("dupe frequency: ", result)
}

func sumInputTwo() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	total, err := CalcTotalWithReader(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("second iteration:", total)
}

// CalcTotal iterates through the input and sums all values
func CalcTotal(input ...string) (int, error) {
	var total int

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

// CalcTotalWithReader sums all values using a reader
func CalcTotalWithReader(r io.Reader) (int, error) {
	var sum int

	s := bufio.NewScanner(r)

	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			return 0, err
		}

		sum += val
	}

	return sum, nil
}

// DuplicateFreq finds the first duplicate frequency calculation
func DuplicateFreq(input []string) int {
	var (
		sum  int
		seen = map[int]bool{}
	)

	for i := 0; i < 1000; i++ {
		for _, i := range input {
			if i == "" {
				continue
			}

			val, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}

			sum += val
			if _, ok := seen[sum]; !ok {
				seen[sum] = true
			} else {
				return sum
			}
		}
	}

	return 0
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
