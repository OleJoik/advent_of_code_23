package main

func main() {
}

type Race struct {
    Time int
    Distance int
}

func task1(races []Race) int {
    output := 1
    for _, r := range races {
        wins := 0
        for speed := 0; speed <= r.Time; speed++ {
            moveTime := r.Time - speed
            distance := moveTime * speed
            if distance > r.Distance {
                wins ++
            }
        }

        output *= wins
    }
    return output
}
