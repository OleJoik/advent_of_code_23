package day2

import (
   "fmt"
   "log"
   "os"
   "bufio"
)

func task1() {
   file, err := os.Open("actual_games.txt")
   if err != nil {
      log.Fatal(err)
   }

   defer file.Close()

   var maximumCubes CubeSet = CubeSet{12, 13, 14}

   scanner := bufio.NewScanner(file)
   sum := 0
   for scanner.Scan() {
   textRow := scanner.Text()
      game := ParseRow(textRow)
      valid := IsValidGame(game, maximumCubes)
      if valid {
         sum = sum + game.Id
      }
   }

   fmt.Println("Sum is", sum)
}


func task2(fileName string) int {
   file, err := os.Open(fileName)
   if err != nil {
      log.Fatal(err)
   }

   defer file.Close()

   scanner := bufio.NewScanner(file)
   sum := 0
   for scanner.Scan() {
   textRow := scanner.Text()
      game := ParseRow(textRow)
      cubeSet := LowestNumberOfCubes(game)
      power := cubeSet.Blue * cubeSet.Red * cubeSet.Green
      
      sum = sum + power
   }

   return sum
}


