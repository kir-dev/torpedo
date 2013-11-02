package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	nums := []float64{1, 2, 3, 4, 5}
	if average(nums) != 3 {
		t.Errorf("Average calculation is wrong for %v", nums)
	}
}

func TestAverageForEmptySlice(t *testing.T) {
	nums := make([]float64, 0)
	actual := average(nums)
	if actual != 0 {
		t.Errorf("Average of empty slice should be 0, got: %f", actual)
	}
}

func TestRemoveFromIntSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := removeInt(slice, 3)
	expected := []int{1, 2, 4, 5}
	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, result)
		}
	}
}

func TestEqualIntSlices(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	if !eqInt(a, b) {
		t.Errorf("Slices (%v and %v) should be equal.", a, b)
	}
}

func TestNotEqualIntSlices(t *testing.T) {
	a := []int{2, 2, 3}
	b := []int{1, 2, 3}
	if eqInt(a, b) {
		t.Errorf("Slices (%v and %v) should not be equal.", a, b)
	}
}
