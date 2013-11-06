package main

var (
	baseShipScore = map[int]float64{
		2: 10,
		3: 8,
		4: 6,
		5: 3,
	}
	// basic deployment
	baseDeployment = []int{5, 4, 3, 3, 2}
)

// basic fleet is [5:1, 4:1, 3:2, 2:1]
const BASE_SCORE = /*5*/ 1*3 + /*4*/ 1*6 + /*3*/ 2*8 + /*2*/ 1*10

type Ship struct {
	Parts  []*ShipPart
	Player *Player
}

type ShipPart struct {
	Field *Field
	IsHit bool
	Ship  *Ship
}

// Creates a new ship with the specfied number of parts. Every part is empty,
// needs further initialization.
func newShip(numberOfParts int) *Ship {
	ship := Ship{}

	ship.Parts = make([]*ShipPart, numberOfParts)
	for i := 0; i < numberOfParts; i++ {
		ship.Parts[i] = &ShipPart{IsHit: false, Ship: &ship}
	}
	return &ship
}

// Determines whether the ship has sunken or not
func (s Ship) isSunken() bool {
	sunken := true
	for _, part := range s.Parts {
		if !part.IsHit {
			sunken = false
			break
		}
	}

	return sunken
}

// Gets the current score of the ship. It is determined by the length of the
// ship and the remaining number of parts and also a base score for the specific
// ship type.
//
// The algorithm is the follow:
//
//   base_score * fitness
func (s Ship) getScore() float64 {
	numberOfParts := len(s.Parts)
	numberOfRemainingParts := 0
	for _, part := range s.Parts {
		if !part.IsHit {
			numberOfRemainingParts++
		}
	}

	fitness := (float64(numberOfRemainingParts) / float64(numberOfParts))
	return baseShipScore[numberOfParts] * fitness
}
