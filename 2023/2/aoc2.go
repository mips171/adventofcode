package main

import (
	"bufio"
	"errors"
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

func checkNextN(fileScanner bufio.Scanner, i int, numTextMap map[string]int, forwards bool) (int, error) {
	for j := 3; j < 6; j++ {
		var nextn string
		if forwards {
			if j > len(fileScanner.Text()[i:]) {
				continue
			}
			nextn = fileScanner.Text()[i:i+j]
		} else {
			if i-j < 0 {
				continue
			}
			nextn = fileScanner.Text()[i-j:i]
		}
		fmt.Println("checking next n: ", nextn)
		if val, ok := numTextMap[nextn]; ok {
			fmt.Println("found a word: ", nextn)
			return val, nil
		}
	}

	return 0, errors.New("no word found")
}

func main() {
    // slice to store the numbers
    numbers := make([]Number, 0)

	// this is pretty small, so we'll just do a map on the stack.
	// but if we had to do combinations, we'd allocate on heap
	numTextMap := make(map[string]int)
	numTextMap["one"] = 1
	numTextMap["two"] = 2
	numTextMap["three"] = 3
	numTextMap["four"] = 4
	numTextMap["five"] = 5
	numTextMap["six"] = 6
	numTextMap["seven"] = 7
	numTextMap["eight"] = 8
	numTextMap["nine"] = 9

    // read input.txt
    // for each line, get the first and last number
    // make a Number struct with those numbers
    // add the Number struct to the slice

    readFile, err := os.Open("aoc2.txt")
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
            if fileScanner.Text()[i] < '0' || fileScanner.Text()[i] > '9' {
                // check if it's a value in our map
				// grab the next n (from 3 to 5) characters, see if any of our words are in it
				// if so, use the value from the map
				n, err := checkNextN(*fileScanner, i, numTextMap, true)
				if err != nil {
					continue
				}
				v1 = n
				break
			}
			if v1 != 0 {
				break
			}
            fmt.Println("first number of the line: ", string(fileScanner.Text()[i]))
            v1, _ = strconv.Atoi(string(fileScanner.Text()[i]))
            break
        }

        // backwards
        for i := len(fileScanner.Text()) - 1; i >= 0; i-- {
            if fileScanner.Text()[i] < '0' || fileScanner.Text()[i] > '9' {
				n, err := checkNextN(*fileScanner, i, numTextMap, true)
				if err != nil {
					continue
				}
				v1 = n
				break
            }
			// we are lucky there are no 'zero's in our search text
			if v2 != 0 {
				break
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
