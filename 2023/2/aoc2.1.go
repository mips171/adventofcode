package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

type PossibleGame struct {
	id int
}

func main() {
	readFile, err := os.Open("aoc2.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	PossibleGamesSet := map[PossibleGame]bool{}

	regexGameId := regexp.MustCompile(`Game (\d+):`)

	for fileScanner.Scan() {
		// split the line by semicolon
		splitByGameMatch := strings.Split(fileScanner.Text(), ";")

		regexRed := regexp.MustCompile(`(\d+) red`)
		regexBlue := regexp.MustCompile(`(\d+) blue`)
		regexGreen := regexp.MustCompile(`(\d+) green`)

		redPossible := true
		bluePossible := true
		greenPossible := true

		for _, splitSemi := range splitByGameMatch {
			fmt.Println("game match: ", splitSemi)

			numRed := regexRed.FindStringSubmatch(splitSemi)
			numBlue := regexBlue.FindStringSubmatch(splitSemi)
			numGreen := regexGreen.FindStringSubmatch(splitSemi)

			if len(numRed) > 0 {
				foundRed, _ := strconv.Atoi(numRed[1])
				if foundRed > MAX_RED {
					fmt.Println("foundRed makes game impossible: ", foundRed)
					redPossible = false
					break
				}
			}

			if len(numBlue) > 0 {
				foundBlue, _ := strconv.Atoi(numBlue[1])
				if foundBlue > MAX_BLUE {
					fmt.Println("foundBlue makes game impossible: ", foundBlue)
					bluePossible = false
					break
				}
			}

			if len(numGreen) > 0 {
				foundGreen, _ := strconv.Atoi(numGreen[1])
				if foundGreen > MAX_GREEN {
					fmt.Println("foundGreen makes game impossible: ", foundGreen)
					greenPossible = false
					break
				}
			}
		}

		if redPossible && bluePossible && greenPossible {
			gameId := regexGameId.FindStringSubmatch(fileScanner.Text())
			gameIdint, _ := strconv.Atoi(strings.Trim(gameId[1], " "))
			fmt.Println("game is possible:", strings.Trim(gameId[1], " "))
			_, ok := PossibleGamesSet[PossibleGame{gameIdint}]
			if !ok {
				PossibleGamesSet[PossibleGame{gameIdint}] = true
			}
		}

		sum := 0
		for g := range PossibleGamesSet {
			sum += g.id
		}
		fmt.Println("sum: ", sum)
	}
}
