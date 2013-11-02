package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
)

var _ = rand.ExpFloat64 //TODO: delete before commit
var _ = errors.New      //TODO: delete before commit

const (
	SIZE = 26
)

const (
	ROW direction = iota
	COLUMN
)

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
}

// Places the appropriate number of ships on the board randomly. The number of
// ships is based on the current average score for the players in the game.
func (board *Board) placeShips(player *Player) {
	scores := make([]float64, len(currentGame.Players))
	for idx, p := range currentGame.Players {
		scores[idx] = p.getCurrentScore()
	}

	avg := average(scores)
	if len(currentGame.Players) == 0 {
		avg = BASE_SCORE
	}

	deployment := computeShipDeployment(avg)
	board.deployShips(player, deployment)
}

func (board *Board) deployShips(player *Player, deployment []int) error {
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
			msg := fmt.Sprintf("number of selected fields (%d) does not match the number of ship parts (%d)",
				len(fields), len(ship.Parts))
			log.Println("ERROR: " + msg)
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
	if isDev() {
		log.Printf("Picking fields [%d, %d] in %s", slot.start+offset, slot.end+offset, dir.toString())
	}

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
			if field.isFree() {
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
