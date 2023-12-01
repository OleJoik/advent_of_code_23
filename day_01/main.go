package main

import (
    "bufio"
    "os"
    "fmt"
    "log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()


	searchStrings := [] string {
		"one", 
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for i := 1; i <= 9; i++ {
		searchStrings = append(searchStrings, strconv.Itoa(i))
    }


	var sum int = 0
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
		firstMatch := FindFirstSubstring(line, searchStrings)
		lastMatch := FindLastSubstring(line, searchStrings)

		joined := ConvertTextToNumber(firstMatch) + ConvertTextToNumber(lastMatch)

		num, _ := strconv.Atoi(joined)

		sum = sum + num
		fmt.Println(line, joined, sum)
    }
}


func FindRegexMatch(pattern string, text string) string {
	spelledOutToNumber := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(text)[1]
	firstNum, ok := spelledOutToNumber[match]
    if !ok {
        firstNum, _ = strconv.Atoi(match)
    }

	return strconv.Itoa(firstNum)
}

func ConvertTextToNumber(text string) string {
	spelledOutToNumber := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	numStr, ok := spelledOutToNumber[text]
    if !ok {
        numStr, _ = strconv.Atoi(text)
    }

	return strconv.Itoa(numStr)
}

func FindFirstSubstring(input string, substrings []string) string {
	minIndex := len(input)
    var firstMatch string

	for _, sub := range substrings {
		index := strings.Index(input, sub)
		if index != -1 && index < minIndex {
            minIndex = index
            firstMatch = sub
        }
	}

	return firstMatch
}


func FindLastSubstring(input string, substrings []string) string {
	maxIndex := 0
    var lastMatch string

	for _, sub := range substrings {
		index := strings.LastIndex(input, sub)
		if index == -1 {
			continue
		}

		fmt.Println("last", sub, index)

		if index >= maxIndex {
            maxIndex = index
            lastMatch = sub
        }
	}

	return lastMatch
}
