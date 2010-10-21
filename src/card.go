package main

type Color int

var colorName = []string{
	"schwarz", "weiß",
	"blau", "grün",
	"gelb", "rot",
}

const (
	Schwarz Color = iota
	Weiß
	Blau
	Grün
	Gelb
	Rot
	deckSize = 48
)

type Deck [deckSize]Card

type Card struct {
	name  string
	color Color
	owner int
}

func (c Color) String() string {
	if c >= Schwarz && c <= Rot {
		return colorName[c]
	}

	return "unbekannt"
}

func (c *Card) String() string {
	return c.name + " (" + c.color.String() + ")"
}
