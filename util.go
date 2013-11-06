package main

import (
	"crypto/rand"
	"errors"
	"fmt"
)

type errorView struct {
	Title   string
	Message string
	IsDev   bool
}

func average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	return sum(numbers) / float64(len(numbers))
}

func sum(nums []float64) float64 {
	var sum float64 = 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

// Removes the first occurence.
func removeInt(slice []int, item int) []int {
	index := -1
	for i, v := range slice {
		if v == item {
			index = i
			break
		}
	}

	if index < 0 {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func eqInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func errorf(format string, vars ...interface{}) error {
	return errors.New(fmt.Sprintf(format, vars...))
}

func generateId() string {
	id := make([]byte, 20)
	rand.Read(id)
	return fmt.Sprintf("%x", id)
}

// Row-column to string
func rcToS(row, col int) string {
	return fmt.Sprintf("(%s, %d)", string(col+65), row+1)
}
