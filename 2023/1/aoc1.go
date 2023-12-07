package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	v1 int
	v2 int
}

func (n Number) String() string {
	return fmt.Sprintf("%d%d", n.v1, n.v2)
}

func main() {
	// slice to store the numbers
	numbers := make([]Number, 0)

	// read input.txt
	// for each line, get the first and last number
	// make a Number struct with those numbers
	// add the Number struct to the slice

    readFile, err := os.Open("aoc1.txt")
    defer readFile.Close()

	if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())

		var v1, v2 int
		// iterate over fileScanner.Text() forwards
		for i := 0; i < len(fileScanner.Text()); i++ {
			if (fileScanner.Text()[i] < '0' || fileScanner.Text()[i] > '9') {
				continue
			}
			fmt.Println("first number of the line: ", string(fileScanner.Text()[i]))
			v1, _ = strconv.Atoi(string(fileScanner.Text()[i]))
			break
		}
		// backwards
		for i := len(fileScanner.Text()) - 1; i >= 0; i-- {
			if (fileScanner.Text()[i] < '0' || fileScanner.Text()[i] > '9') {
				continue
			}
			fmt.Println("last number of the line: ", string(fileScanner.Text()[i]))
			v2, _ = strconv.Atoi(string(fileScanner.Text()[i]))
			break
		}

		numbers = append(numbers, Number{v1, v2})
	}

	fullSum := 0

	for _, number := range numbers {
		num, err := strconv.Atoi(number.String())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fullSum += num
	}

	fmt.Println(fullSum)
}