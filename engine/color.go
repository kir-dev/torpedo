package engine

type Color struct {
	Hit        string
	HitAndSunk string
}

var (
	colors []Color = []Color{
		Color{"#FF0000", "#FFB5B5"}, // red
		Color{"#008108", "#89C18C"}, // green
		Color{"#002F81", "#B2C3E1"}, // blue
		Color{"#FFF500", "#FFF5A2"}, // yellow
		Color{"#8C0062", "#FFC3ED"}, // purple
		Color{"#FF7800", "#FFBE83"}, // orange
		Color{"#000000", "#989898"}, // black
		Color{"#00AB83", "#C0EBE1"}, // turkiz
		Color{"#5A4F3C", "#AA946C"}, // brown
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
