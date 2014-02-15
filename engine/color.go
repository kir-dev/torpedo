package engine

type Color struct {
	Hit        string
	HitAndSunk string
}

var (
	colors []Color = []Color{
		Color{"blue", "purple"},
		Color{"yellow", "black"},
		Color{"olive", "tomato"},
	}
	counter = 0
)

func getNextColor() Color {
	c := colors[counter]
	counter += 1
	return c
}

func ResetColors() {
	counter = 0
}
