package day2

import (
	"strconv"
	"strings"
)

type Game struct {
   Id int
   Sets []CubeSet
}

type CubeSet struct {
   Red int
   Green int
   Blue int
}

func ParseRow(row string) Game {
   stringElements := strings.Split(row, ":")
   gameId, err := strconv.Atoi(strings.Split(stringElements[0], " ")[1])
   if err != nil {
      panic(err)
   }
   
   cubeSetStrings := strings.Split(stringElements[1], ";")
   cubeSets := []CubeSet {}

   for _, cubeSetString := range cubeSetStrings {
      cubeSetColorsString := strings.Split(cubeSetString, ",")

      cubeSet := CubeSet{Red: 0, Green: 0, Blue: 0,}

      for _, cubeSetColorString := range cubeSetColorsString {
         trimmedString := strings.Trim(cubeSetColorString, " ")
         splitCubeSetColorString := strings.Split(trimmedString, " ")
         count := splitCubeSetColorString[0]
         color := splitCubeSetColorString[1]

         switch color {
            case "red":
                cubeSet.Red, err = strconv.Atoi(count)
            case "green":
                cubeSet.Green, err = strconv.Atoi(count)
            case "blue":
                cubeSet.Blue, err = strconv.Atoi(count)
            default:
                panic("Did not recognize the color")
         }

         if err != nil {
              panic("Did not recognize the color")
         }
      }

      cubeSets = append(cubeSets, cubeSet)
   }



   return Game{Id: gameId, Sets: cubeSets}
}

func IsValidGame(game Game, maximumCubes CubeSet) bool {
   for _, set := range game.Sets {
      if set.Red > maximumCubes.Red {
         return false
      }

      if set.Green > maximumCubes.Green {
         return false
      }

      if set.Blue > maximumCubes.Blue {
         return false
      }
   }

   return true
}


func LowestNumberOfCubes(game Game) CubeSet {
   minimum := CubeSet{Red: 0, Green: 0, Blue: 0, }

   for _, set := range game.Sets {
      if set.Red > minimum.Red {
         minimum.Red = set.Red
      }

      if set.Green > minimum.Green {
         minimum.Green = set.Green
      }

      if set.Blue > minimum.Blue {
         minimum.Blue = set.Blue
      }
   }

   return minimum
}
