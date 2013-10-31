package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
    values := []int{123, 3064, 3, 64, 490}
    BubbleSort(values)
    if values[0] != 3 || values[1] != 64 || values[2] != 123 || values[3] != 490 || values[4] != 3064 {
        t.Error("Bubblesort() failed. Got", values, "Expected 3 64 123 490 3064")
    }   
}

func TestBubbleSort2(t *testing.T) {
    values := []int{5, 5, 3, 2, 1}
    BubbleSort(values)
    if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
        t.Error("Bubblesort() failed. Got", values, "Expected 1 2 3 5 5")
    }   
}

func TestBubbleSort3(t *testing.T) {
    values := []int{5}
    BubbleSort(values)
    if values[0] != 5 {
        t.Error("Bubblesort() failed. Got", values, "Expected 5")
    }   
}
