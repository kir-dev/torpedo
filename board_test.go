package main

import (
	"testing"
)

func TestComputeDeployments(t *testing.T) {
	var deploymentTests = []struct {
		avg      float64
		expected []int
	}{
		{float64(BASE_SCORE), baseDeployment},
		{float64(BASE_SCORE) - baseShipScore[5], []int{4, 3, 3, 2}},
		{BASE_SCORE / 2.0, []int{5, 4, 3}},
		{0, []int{}},
	}

	for _, testCase := range deploymentTests {
		depl := computeShipDeployment(testCase.avg)

		if !eqInt(depl, testCase.expected) {
			t.Errorf("Wrong deployment for %v average score. Expected: %v, got: %v", testCase.avg, testCase.expected, depl)
		}
	}
}

func TestGetColumn(t *testing.T) {
	board := newGame().Board
	fields := board.getColumn(1, 1, 3)

	if len(fields) != 2 {
		t.Error("Wrong number of fields selected for column.")
	}
}
