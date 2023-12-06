package main

import "testing"


func TestTask1Example(t *testing.T) {
    races := []Race {
        {Time: 7, Distance: 9}, 
        {Time: 15, Distance: 40},
        {Time: 30, Distance: 200},
    }
    output := task1(races)

    expected := 288

    if output != expected {
        t.Error("Output is", output, "should be", expected)
    }
}

func TestTask1Actual(t *testing.T) {
    races := []Race {
        {Time: 58, Distance: 434}, 
        {Time: 81, Distance: 1041},
        {Time: 96, Distance: 2219},
        {Time: 76, Distance: 1218},
    }

    output := task1(races)

    expected := 1159152

    if output != expected {
        t.Error("Output is", output, "should be", expected)
    }
}

func TestTask2Example(t *testing.T) {
    races := []Race {
        {Time: 71530, Distance: 940200}, 
    }

    output := task1(races)

    expected := 71503

    if output != expected {
        t.Error("Output is", output, "should be", expected)
    }
}


func TestTask2Actual(t *testing.T) {
    races := []Race {
        {Time: 58819676, Distance: 434104122191218}, 
    }

    output := task1(races)

    expected := 71503

    if output != expected {
        t.Error("Output is", output, "should be", expected)
    }
}
