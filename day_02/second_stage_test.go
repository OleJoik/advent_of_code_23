package day2

import "testing"

func TestTask2Example(t *testing.T) {
   value := task2("example_games.txt")

   if value != 2286 {
        t.Error("Value is not correct, should be 2286, but is", value)
   }

}

func TestTask2Actual(t *testing.T) {
   value := task2("actual_games.txt")

   if value != 84911 {
        t.Error("Value is not correct, should be 84911, but is", value)
   }
}
