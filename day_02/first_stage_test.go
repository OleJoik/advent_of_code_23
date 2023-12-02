package day2

import "testing"

func TestGame1(t *testing.T) {
    var maximumCubes CubeSet = CubeSet{ Red: 12, Green: 13, Blue: 14, }

    text := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

    game := ParseRow(text)

    valid := IsValidGame(game, maximumCubes)

    if !valid {
        t.Errorf("Game 1 should be valid")
    }
}

func TestGame2(t *testing.T) {
    var maximumCubes CubeSet = CubeSet{ Red: 12, Green: 13, Blue: 14, }

    text := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"

    game := ParseRow(text)

    valid := IsValidGame(game, maximumCubes)

    if !valid {
        t.Errorf("Game 2 should be valid")
    }
}

func TestGame3(t *testing.T) {
    var maximumCubes CubeSet = CubeSet{ Red: 12, Green: 13, Blue: 14, }

    text := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"

    game := ParseRow(text)

    valid := IsValidGame(game, maximumCubes)

    if valid {
        t.Errorf("Game 2 should be INVALID")
    }
}

func TestGame4(t *testing.T) {
    var maximumCubes CubeSet = CubeSet{ Red: 12, Green: 13, Blue: 14, }

    text := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"

    game := ParseRow(text)

    valid := IsValidGame(game, maximumCubes)

    if valid {
        t.Errorf("Game 4 should be INVALID")
    }
}


func TestGame5(t *testing.T) {
    var maximumCubes CubeSet = CubeSet{ Red: 12, Green: 13, Blue: 14, }

    text := "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"

    game := ParseRow(text)

    valid := IsValidGame(game, maximumCubes)

    if !valid {
        t.Errorf("Game 5 should be valid")
    }
}

func TestGame96(t *testing.T) {
    var maximumCubes CubeSet = CubeSet{ Red: 12, Green: 13, Blue: 14, }
    text := "Game 96: 5 red, 7 green; 4 blue, 14 green, 10 red; 13 green; 13 green, 3 blue; 13 green, 1 red, 3 blue; 12 red, 1 green"

    game := ParseRow(text)

    valid := IsValidGame(game, maximumCubes)

    if valid {
        t.Errorf("Game 96 should be INVALID")
    }
}

func TestTask1(t *testing.T) {
    task1()
}
