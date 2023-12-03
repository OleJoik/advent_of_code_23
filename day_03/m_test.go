package main

import (
	"testing"
)

func TestTask1Example(t *testing.T) {
    actual := task1("input_example.txt")

    expected := 4361
    if actual != expected {
        t.Errorf("Value is not correct, should be %d, but is %d", expected, actual)
    }
}

func TestTask1Actual(t *testing.T) {
    actual := task1("input_actual.txt")

    expected := 525181
    if actual != expected {
        t.Errorf("Value is not correct, should be %d, but is %d", expected, actual)
    }
}


func TestParseError1(t *testing.T) {
    row := "..*392.423$."

    newNumbers := FindNumberData(0, row)

    actual := newNumbers[1].Number

    expected := 423
    if actual != expected {
        t.Errorf("Value is not correct, should be %d, but is %d", expected, actual)
    }
}

func TestTask2Example(t *testing.T) {
    actual := task2("input_example.txt")

    expected := 467835
    if actual != expected {
        t.Errorf("Value is not correct, should be %d, but is %d", expected, actual)
    }
}

func TestTask2Actual(t *testing.T) {
    actual := task2("input_actual.txt")

    expected := 84289137
    if actual != expected {
        t.Errorf("Value is not correct, should be more than %d, but is %d", expected, actual)
    }
}
