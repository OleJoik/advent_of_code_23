package main

import "testing"

func TestTask1Actual(t *testing.T) {
    output := task1("data_actual.txt")

    expected := 993500720

    if output != expected {
        t.Error("Output is", output, "should be", expected)
    }
}


func TestTask2Example(t *testing.T) {
    output := task2("data_example.txt")

    expected := 46

    if output != expected {
        t.Error("Output is", output, "should be", expected)
    }
}


func TestTask2Actual(t *testing.T) {
    output := task2("data_actual.txt")

    expected := 4917124

    if output != expected {
        t.Error("Output is", output, "should be", expected)
    }
}
