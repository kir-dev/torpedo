package main

import (
	"bytes"
	"github.com/kir-dev/torpedo/engine"
	"github.com/kir-dev/torpedo/util"
	"html/template"
)

const (
	CELL_BASE_COLOR   = "white"
	CELL_MISS_COLOR   = "gray"
	CELL_PLAYER_COLOR = "lime"
)

func utilFuncMap() template.FuncMap {
	return template.FuncMap{
		"add":           add,
		"letters":       letters,
		"hasWinner":     hasWinner,
		"ship_color":    getShipColor,
		"ship_color2":   getShipColorOnShootPage,
		"num_to_letter": getLetterForNumber,
	}
}

func hasWinner(winner *engine.Player) bool {
	if winner != nil {
		return true
	} else {
		return false
	}
}

func add(a, b int) int {
	return a + b
}

func letters(count int, tag string) template.HTML {
	t, err := template.New("partail").Parse("<" + tag + ">{{.}}</" + tag + ">")

	if err != nil {
		util.LogError(err.Error())
		return ""
	}

	buffer := bytes.Buffer{}

	for i := -1; i < count; i++ {
		if i < 0 {
			t.Execute(&buffer, nil)
		} else {
			t.Execute(&buffer, getLetterForNumber(i))
		}
	}

	return template.HTML(buffer.String())
}

func getShipColor(field *engine.Field) string {
	if field.IsHit {
		if field.IsEmpty() {
			return CELL_MISS_COLOR
		}
		if field.ShipPart.Ship.IsSunken() {
			return field.ShipPart.Ship.Player.Color.HitAndSunk
		}
		return field.ShipPart.Ship.Player.Color.Hit
	}

	return CELL_BASE_COLOR
}

func getShipColorOnShootPage(field *engine.Field, player *engine.Player) string {
	if !field.IsHit && field.GetPlayerId() == player.Id {
		return CELL_PLAYER_COLOR
	}
	return getShipColor(field)
}

func getLetterForNumber(num int) string {
	return string(num + 65)
}
