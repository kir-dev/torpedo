package engine

type gap struct {
	start int
	end   int
}

func (s gap) size() int {
	return absInt(s.end-s.start) + 1
}

func findEmptySlot(size int, fields []*Field) (gap, bool) {
	slot := gap{-1, -1}
	for i, f := range fields {
		if slot.size() == size {
			break
		}

		if f.IsEmpty() {
			if slot.start < 0 {
				slot.start = i
				slot.end = i
			} else {
				slot.end = i
			}
		} else {
			slot = gap{-1, -1}
		}
	}

	if slot.size() == size {
		return slot, true
	}

	return slot, false
}

func eqShipSlot(a, b gap) bool {
	return a.start == b.start && a.end == b.end
}
