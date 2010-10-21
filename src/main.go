package main

import "fmt"

const (
	deckfile = "../cards"
)

func main() {
	openend := 0

	fmt.Printf("Laden...\n")
	deck := readDeckFile(deckfile)

	printTable(deck, -1, "")

	for openend < deckSize {

		for valid, c := false, promptCard(); !valid; {
			if deck[c].owner == 0 {
				valid = true
				fmt.Println()
				printTable(deck, c, "")
				fmt.Printf("Karte %s eingesammelt\n\n", deck[c].String())
				deck[c].owner++
				openend++
			} else {
				fmt.Printf("Karte %d ist schon eingesammelt\n", c + 1)
				c = promptCard()
			}
		}
	}
}
