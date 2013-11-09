package engine

import (
	"testing"
)

func TestFindEmptySlot(t *testing.T) {
	var testdata = []struct {
		size     int
		fields   []*Field
		expected gap
		ok       bool
	}{
		{
			2,
			[]*Field{
				&Field{},
				&Field{},
			},
			gap{0, 1},
			true,
		},
		{
			2,
			[]*Field{
				&Field{ShipPart: &ShipPart{}},
				&Field{},
				&Field{},
			},
			gap{1, 2},
			true,
		},
		{
			2,
			[]*Field{
				&Field{ShipPart: &ShipPart{}},
				&Field{ShipPart: &ShipPart{}},
				&Field{ShipPart: &ShipPart{}},
			},
			gap{-1, -1},
			false,
		},
		{
			3,
			[]*Field{
				&Field{ShipPart: &ShipPart{}},
				&Field{},
				&Field{},
				&Field{ShipPart: &ShipPart{}},
			},
			gap{-1, -1},
			false,
		},
	}

	for i, td := range testdata {
		actual, ok := findEmptySlot(td.size, td.fields)
		if !eqShipSlot(td.expected, actual) {
			t.Errorf("Expected ship slot: %v, got: %v in test data: %d", td.expected, actual, i)
		}

		if ok != td.ok {
			t.Errorf("In %d example the expected: %v, got: %v", i, td.ok, ok)
		}
	}

}
