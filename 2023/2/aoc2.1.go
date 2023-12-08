package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PossibleGamePowerSet struct {
	r int
	g int
	b int
}

func main() {
	readFile, err := os.Open("aoc2.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	PossibleGames := make([]PossibleGamePowerSet, 0)

	for fileScanner.Scan() {
		splitByGameMatch := strings.Split(fileScanner.Text(), ";")

		regexRed := regexp.MustCompile(`(\d+) red`)
		regexBlue := regexp.MustCompile(`(\d+) blue`)
		regexGreen := regexp.MustCompile(`(\d+) green`)

		redPossible := true
		bluePossible := true
		greenPossible := true

		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		for _, splitSemi := range splitByGameMatch {
			fmt.Println("game match:", splitSemi)

			numRed := regexRed.FindStringSubmatch(splitSemi)
			numBlue := regexBlue.FindStringSubmatch(splitSemi)
			numGreen := regexGreen.FindStringSubmatch(splitSemi)

			if len(numRed) > 0 {
				foundRed, _ := strconv.Atoi(numRed[1])
				if (foundRed > maxRed) && redPossible {
					fmt.Println("found a new max red:", foundRed)
					maxRed = foundRed
				}
			}

			if len(numBlue) > 0 {
				foundBlue, _ := strconv.Atoi(numBlue[1])
				if (foundBlue > maxBlue) && bluePossible {
					maxBlue = foundBlue
				}
			}

			if len(numGreen) > 0 {
				foundGreen, _ := strconv.Atoi(numGreen[1])
				if foundGreen > maxGreen && greenPossible {
					maxGreen = foundGreen
				}
			}
		}

		if redPossible && bluePossible && greenPossible {
			PossibleGames = append(PossibleGames, PossibleGamePowerSet{maxRed, maxBlue, maxGreen})
		}
	}

	sum := 0
	for _, g := range PossibleGames {
		sum += g.r * g.g * g.b
	}
	fmt.Println("sum:", sum)
}
