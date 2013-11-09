package engine

import (
	"errors"
	"fmt"
	"github.com/kir-dev/torpedo/util"
	"math"
	"math/rand"
	"strconv"
	"sync"
)

const (
	SIZE = 26
)

const (
	ROW direction = iota
	COLUMN
)

const (
	HIT      = HitResult("hit")
	HIT_SUNK = HitResult("hit&sunk")
	MISS     = HitResult("miss")
	INVALID  = HitResult("invalid")
)

type HitResult string

type direction int

func (d direction) toString() string {
	switch {
	case d == ROW:
		return "ROW"
	case d == COLUMN:
		return "COLUMN"
	}

	return "unknown"
}

// Represents the board of the game
type Board struct {
	Fields [SIZE][SIZE]*Field
	mu     sync.Mutex
}

// Places the appropriate number of ships on the board randomly. The number of
// ships is based on the current average score for the players in the game.
func (board *Board) placeShips(allPlayers []*Player, player *Player) error {
	scores := make([]float64, len(allPlayers))
	for idx, p := range allPlayers {
		scores[idx] = p.getCurrentScore()
	}

	avg := average(scores)
	util.LogInfo("average board score is %f", avg)
	if len(allPlayers) == 0 {
		avg = BASE_SCORE
	}

	deployment := computeShipDeployment(avg)
	util.LogInfo("Deployment for %s player: %v", player.Name, deployment)
	return board.deployShips(player, deployment)
}

// Deploy ships to on board for the given player.
func (board *Board) deployShips(player *Player, deployment []int) error {
	// lock this method. only one player can deploy ships at once
	board.mu.Lock()
	defer board.mu.Unlock()

	ships := make([]*Ship, len(deployment))
	for idx, size := range deployment {
		ship := newShip(size)
		row := rand.Intn(SIZE)
		col := rand.Intn(SIZE)

		// TODO: introduce max retry count?
		fields, err := board.chooseFields(size, row, col)
		for err != nil {
			row = rand.Intn(SIZE)
			col = rand.Intn(SIZE)
			fields, err = board.chooseFields(size, row, col)
		}

		if len(fields) != len(ship.Parts) {
			msg := fmt.Sprintf(
				"Number of selected fields (%d) does not match the number of ship parts (%d)",
				len(fields),
				len(ship.Parts),
			)
			return errors.New(msg)
		}

		// build connection between field and ship part
		for i := 0; i < len(fields); i++ {
			fields[i].ShipPart = ship.Parts[i]
			ship.Parts[i].Field = fields[i]
		}
		ships[idx] = ship
	}
	player.Ships = ships
	return nil
}

// choose fields for ship
func (board *Board) chooseFields(size, row, col int) ([]*Field, error) {
	// detect if it could be fit into the row
	start_row := maxInt(col-(size-1), 0)
	end_row := minInt(col+(size-1), SIZE-1)
	slot_row, ok_row := findEmptySlot(size, board.Fields[row][start_row:end_row])

	// detect if it could be fit into the column
	start_col := maxInt(row-(size-1), 0)
	end_col := minInt(row+(size-1), SIZE-1)
	slot_col, ok_col := findEmptySlot(size, board.getColumn(col, start_col, end_col))

	bothOk := ok_col && ok_row
	randDir := rand.Intn(2)

	// return something.
	switch {
	case (ok_row && !ok_col) || (bothOk && randDir == int(ROW)):
		return board.fieldsInSlot(row, col, start_row, ROW, slot_row), nil
	case (!ok_row && ok_col) || (bothOk && randDir == int(COLUMN)):
		return board.fieldsInSlot(row, col, start_col, COLUMN, slot_col), nil
	}

	// no fit, return error
	return nil, errors.New(fmt.Sprintf("No suitable place slot for %d size ship on (%d, %d)", size, row, col))
}

// Slice in a specific direction
func (b *Board) fieldsInSlot(row, col, offset int, dir direction, slot gap) []*Field {
	util.LogDebug("Picking fields [%d, %d] in %s", slot.start+offset, slot.end+offset, dir.toString())
	if dir == ROW {
		return b.Fields[row][slot.start+offset : slot.end+offset+1]
	}
	return b.getColumn(col, slot.start+offset, slot.end+offset+1)
}

// Gets the fields from a column [start, end).
func (b *Board) getColumn(col, start, end int) []*Field {
	var fields []*Field
	// cannot slice on the vertically in the array...
	for i := start; i < end; i++ {
		fields = append(fields, b.Fields[i][col])
	}
	return fields
}

func (board *Board) print() {
	fmt.Print("    ")
	for i := 0; i < SIZE; i++ {
		letter, _ := strconv.Unquote(fmt.Sprintf("%q", i+65))
		fmt.Print(letter)
	}
	fmt.Println()
	for i, row := range board.Fields {
		fmt.Printf("%02d: ", i+1)
		for _, field := range row {
			if field.IsEmpty() {
				fmt.Print("_")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func computeShipDeployment(boardAverage float64) []int {
	diff := BASE_SCORE - boardAverage
	deployment := make([]int, len(baseDeployment))
	copy(deployment, baseDeployment)

	var min float64 = -1
	var sizeToExclude int

	for diff > 2.0 {
		for _, size := range deployment {
			currDiff := math.Abs(diff - baseShipScore[size])
			if min < 0 || currDiff < min {
				min = currDiff
				sizeToExclude = size
			}
		}

		deployment = removeInt(deployment, sizeToExclude)
		diff = diff - baseShipScore[sizeToExclude]
	}

	return deployment
}

// Shoot at a coordinate on the game's board
func (b *Board) shootAt(row, col int, endTurn chan<- int) HitResult {
	field := b.Fields[row][col]

	if field.IsHit {
		return INVALID
	}

	result := MISS
	field.IsHit = true
	if !field.IsEmpty() {
		field.ShipPart.IsHit = true
		if field.ShipPart.Ship.isSunken() {
			result = HIT_SUNK
		} else {
			result = HIT
		}
	}

	// signal the timer that this turn ended
	if endTurn != nil {
		close(endTurn)
	}

	return result
}
