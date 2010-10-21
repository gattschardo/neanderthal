package card

import "testing"

type colorStringTest struct {
	in Color
	out string
}

var colorStringTests = []colorStringTest{
	colorStringTest{Schwarz, "schwarz"},
	colorStringTest{Weiß, "weiß"},
	colorStringTest{Blau, "blau"},
	colorStringTest{Grün, "grün"},
	colorStringTest{Gelb, "gelb"},
	colorStringTest{Rot, "rot"},
}

func TestColorString(t *testing.T) {
	for _, cst := range colorStringTests {
		if s := cst.in.String(); s != cst.out {
			t.Errorf("%s.String() says %s, want %s", cst.in, s, cst.out);
		}
	}
}

type cardStringTest struct {
	in Card
	out string
}

var cardStringTests = []cardStringTest{
	cardStringTest{ Card{"name",Schwarz}, "name (schwarz)"},
	cardStringTest{ Card{"",Schwarz}, " (schwarz)"},
}

func TestCardString(t *testing.T) {
	for _, cst := range cardStringTests {
		if s := cst.in.String(); s != cst.out {
			t.Errorf("%s.String() says %s, want %s", cst.in, s, cst.out);
		}
	}
}
