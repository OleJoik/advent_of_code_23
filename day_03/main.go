package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
    "strconv"
)

func main() {
    task1("input_actual.txt")
}

func task1(filepath string) int {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    rows := []string {}
    for scanner.Scan() {
        rows = append(rows, "." + strings.TrimSpace(scanner.Text()) + ".")
    }

    sum := 0
    numberData := []NumberData {}

    for row, textRow := range rows {
        newNumbers := FindNumberData(row, textRow)
        for _, num := range newNumbers {
            numberData = append(numberData, num)
        }
    }

    for i, d := range numberData {
        left := PotentialGear{
            Symbol: rune(rows[d.Row][d.Start - 1]), 
            Count: 0, 
            Row: d.Row, 
            Column: d.Start - 1,
        }

        right := PotentialGear{
            Symbol: rune(rows[d.Row][d.End + 1]), 
            Count: 0, 
            Row: d.Row, 
            Column: d.Start + 1,
        }

        runes := []PotentialGear {left, right}

        if d.Row != 0 {
            runesAbove := rows[d.Row - 1][d.Start - 1:d.End + 2]

            col := d.Start - 1
            for _, r := range runesAbove {
                runes = append(runes, PotentialGear{
                    Symbol: rune(r), 
                    Count: 0, 
                    Row: d.Row, 
                    Column: col,
                })
                col++
            }
        }

        if d.Row != len(rows) - 1 {
            runesBelow := rows[d.Row + 1][d.Start - 1:d.End + 2]

            col := d.Start - 1
            for _, r := range runesBelow {
                runes = append(runes, PotentialGear{
                    Symbol: rune(r), 
                    Count: 0, 
                    Row: d.Row, 
                    Column: col,
                })
                col++
            }
        }

        allRunes := []rune {}
        for _, g := range runes {
            allRunes = append(allRunes, g.Symbol)
        }

        if HasSymbol(allRunes) {
            numberData[i].IsValid = true
            sum += d.Number
        }
    }

    fmt.Println("sum", sum)


    return sum
}


func task2(filepath string) int {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    rows := []string {}
    for scanner.Scan() {
        rows = append(rows, "." + strings.TrimSpace(scanner.Text()) + ".")
    }

    sum := 0
    numberData := []NumberData {}

    for row, textRow := range rows {
        newNumbers := FindNumberData(row, textRow)
        for _, num := range newNumbers {
            numberData = append(numberData, num)
        }
    }

    potentialGears := []PotentialGear {}

    for _, d := range numberData {
        left := PotentialGear{
            Symbol: rune(rows[d.Row][d.Start - 1]), 
            Count: 0, 
            Row: d.Row, 
            Column: d.Start - 1,
            Numbers: []NumberData { d },
        }

        right := PotentialGear{
            Symbol: rune(rows[d.Row][d.End + 1]), 
            Count: 0, 
            Row: d.Row, 
            Column: d.End + 1,
            Numbers: []NumberData { d },
        }

        potentialGears = append(potentialGears, left)
        potentialGears = append(potentialGears, right)

        if d.Row != 0 {
            runesAbove := rows[d.Row - 1][d.Start - 1:d.End + 2]

            col := d.Start - 1
            for _, r := range runesAbove {
                potentialGears = append(potentialGears, PotentialGear{
                    Symbol: rune(r), 
                    Count: 0, 
                    Row: d.Row - 1, 
                    Column: col,
                    Numbers: []NumberData { d },
                })
                col++
            }
        }

        if d.Row != len(rows) - 1 {
            runesBelow := rows[d.Row + 1][d.Start - 1:d.End + 2]

            col := d.Start - 1
            for _, r := range runesBelow {
                potentialGears = append(potentialGears, PotentialGear{
                    Symbol: rune(r), 
                    Count: 0, 
                    Row: d.Row + 1, 
                    Column: col,
                    Numbers: []NumberData { d },
                })
                col++
            }
        }
    }

    filteredGears := []PotentialGear {}


    for _, g := range potentialGears {
        if string(g.Symbol) == "*" {
            filteredGears = append(filteredGears, g)
        }
    }

    for _, g := range filteredGears {
        result := 1
        displayNumbers:= ""
        for _, n := range g.Numbers {
            result *= n.Number
            displayNumbers += " "
            displayNumbers += fmt.Sprint(n.Number)
        }
    }

    newFiltered := []PotentialGear {}
    for _, g := range filteredGears {
        exists := false
        for j, n := range newFiltered {
            if g.Row == n.Row && g.Column == n.Column {
                exists = true
                newFiltered[j].Count ++
                newFiltered[j].Numbers = append(newFiltered[j].Numbers, g.Numbers[0])
            }
        }

        if !exists {
            g.Count = 1
            newFiltered = append(newFiltered, g)
        }
    }

    filteredGears = newFiltered
    newFiltered = []PotentialGear {}

    for _, g := range filteredGears {
        if g.Count == 2 {
            newFiltered = append(newFiltered, g)
        }
    }

    filteredGears = newFiltered

    for _, g := range filteredGears {
        result := 1
        displayNumbers:= ""
        for _, n := range g.Numbers {
            result *= n.Number
            displayNumbers += " "
            displayNumbers += fmt.Sprint(n.Number)
        }

        sum += result
    }

    
    fmt.Println("sum", sum)


    return sum
}

type NumberData struct {
    Number int
    Row int
    Start int
    End int
    IsValid bool
}

type PotentialGear struct {
    Symbol rune
    Row int
    Column int
    Count int
    Numbers []NumberData
}

func FindNumberData(rowNumber int, row string) []NumberData {
    startNum := 0
    endNum := 0
    isInNum := false

    numbers := []NumberData {}

    for i, char := range row {
        

        if !isInNum && unicode.IsNumber(char) {
            startNum = i
            isInNum = true
        }

        if isInNum && !unicode.IsNumber(char) {
            endNum = i - 1
            isInNum = false
            num, _ := strconv.Atoi(row[startNum:i])
            numbers = append(numbers, NumberData{
                Row: rowNumber, 
                Start: startNum, 
                End: endNum,
                Number: num,
            })
        }
    }

    return numbers
}


func HasSymbol(runes []rune) bool {
    hasSymbol := false
    for _, rune := range runes {
        if string(rune) == "." {
            continue
		} else {
            hasSymbol = true
        }
    }

    return hasSymbol
}
