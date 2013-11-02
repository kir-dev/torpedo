package main

import (
	"testing"
)

func TestShipIsSunken(t *testing.T) {
	s := Ship{
		[]*ShipPart{
			&ShipPart{IsHit: true},
			&ShipPart{IsHit: true},
		},
		nil,
	}

	if !s.isSunken() {
		t.Error("Ship should be sunken.")
	}
}

func TestShipIsSunkenWithOneUnbrokenPart(t *testing.T) {
	s := Ship{
		[]*ShipPart{
			&ShipPart{IsHit: true},
			&ShipPart{IsHit: false},
		},
		nil,
	}

	if s.isSunken() {
		t.Error("Ship should not be sunken.")
	}

}

func TestIntactShipHasBaseScore(t *testing.T) {
	s := newShip(5)

	if s.getScore() != baseShipScore[5] {
		t.Error("Intact ship should have its base score")
	}
}
