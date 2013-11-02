package main

// Represents a field on the board.
type Field struct {
	// Shows if this particural field has been hit or not.
	IsHit bool
	// The ship part that is on this field or nil if there's none.
	ShipPart *ShipPart
}

func (f Field) IsFree() bool {
	return f.ShipPart == nil
}
