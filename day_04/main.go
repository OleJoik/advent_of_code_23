package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("Hello world")
}

func task1(filepath string) int {
    seeds, categoryMap := parseFile(filepath)
    locations := []int {}
    for _, s := range seeds {
        t := "seed"
        v := s

        for t != "location" {
            t, v = remap(t, v, categoryMap)
        }

        locations = append(locations, v)
    }

    sort.Ints(locations)

    return locations[0]
}

type SeedSet struct {
    Start int
    Range int
}

func task2(filepath string) int {
    seeds, categoryMap := parseFile(filepath)

    seedSet := []SeedSet {}   
    for i := 0; i<len(seeds); i += 2 {
        seedSet = append(seedSet, SeedSet{Start: seeds[i], Range: seeds[i+1]})
    }


    totalIterations := len(seedSet) 
    for i, s := range seedSet {
        fmt.Println("Range in seedset", i + 1, s.Range)
    }

    location := 99999999999999999
    for j, set := range seedSet {
        for i := set.Start; i < set.Start + set.Range; i++ {
            t := "seed"
            v := i 

            for t != "location" {
                t, v = remap(t, v, categoryMap)
            }

            if v < location {
                location = v
            }
        }

        fmt.Println("Iterations: ", j + 1, "/", totalIterations )
    }

    fmt.Println("sorting...")

    fmt.Println("done sorting...")

    return location
}




func remap(current string, value int, categoryMap map[string] Map) (string, int) {
    newMap := categoryMap[current]
    toType := newMap.To
    routes := newMap.Routes

    for _, r := range routes {
        if value >= r.SourceStart && value < r.SourceStart + r.Length {
            newVal := r.DestinationStart + value - r.SourceStart

            return toType, newVal
        }
    }

    return toType, value
}

type MapBoundaries struct {
    TitleRow int
    Start int
    End int
}

type Map struct {
    To string
    Routes []Router
}

type Router struct {
    Length int
    SourceStart int
    DestinationStart int
}

type Current struct {
    Type string
    value int
}

func parseFile(filepath string) ([]int, map[string] Map) {
    mapCategories := make(map[string] Map)
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    rows := []string {}
    for scanner.Scan() {
        rows = append(rows, strings.TrimSpace(scanner.Text()))
    }

    seedInput := strings.Split(rows[0], " ")[1:]
    seeds := []int { }
    for _, s := range seedInput {
        seed, _ := strconv.Atoi(s)
        seeds = append(seeds, seed)
    }

    emptyRows := []int {}
    for i, row := range rows {
        if row == "" {
            emptyRows = append(emptyRows, i)
        }
    }

    mapBoundaries := [] MapBoundaries {}

    for i := 0; i<len(emptyRows) - 1; i++ {
        r := emptyRows[i]
        mapBoundaries = append(mapBoundaries, MapBoundaries{TitleRow: r+1, Start: r+2, End: emptyRows[i+1]})
    }

    for _, b := range mapBoundaries {
        categories := strings.Split(strings.Split(rows[b.TitleRow], " ")[0], "-")
        
        routes := [] Router {}

        for i := b.Start; i < b.End; i++ {
            n := strings.Split(rows[i], " ")
            destinationStart, _ := strconv.Atoi(n[0])
            sourceStart, _ := strconv.Atoi(n[1])
            rangeLength, _ := strconv.Atoi(n[2])
            routes = append(routes, Router{SourceStart: sourceStart, DestinationStart: destinationStart, Length: rangeLength})
        }



        mapCategories[categories[0]] = Map{To: categories[2], Routes: routes}
    }


    return seeds, mapCategories
}

