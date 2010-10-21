package main

import (
	"os"
	"fmt"
	"bufio"
)

const (
	deckfile = "../cards"
)

func main() {
	openend := 0

	fmt.Printf("Laden...\n")
	setup()
	in := bufio.NewReader(os.Stdin)
	deck := readDeckFile(deckfile)

	next := needed()
	fmt.Print("\nWurfholz werfen mit [Enter]> ")
	_, err := in.ReadString('\n')
	if err == os.EOF { return }
	printTable(deck, -1, next.String())

	for openend < deckSize {

		for valid, c := false, promptCard(in); !valid; {

			if deck[c].owner == 0 {
				if deck[c].color == next {
					valid = true
					fmt.Println()
					printTable(deck, c, next.String())
					fmt.Printf("Karte %s eingesammelt\n", deck[c].String())
					deck[c].owner++
					openend++
				} else {
					valid = true
					fmt.Println()
					printTable(deck, c, next.String())
					fmt.Printf("Falsche Farbe!\n")
				}
			} else {
				fmt.Printf("Karte %d ist schon eingesammelt\n", c + 1)
				c = promptCard(in)
			}
		}

		next = needed()
		fmt.Print("\nWurfholz werfen mit [Enter]> ")
		s, err := in.ReadString('\n')
		if err == os.EOF || s == "q\n" { return }
		printTable(deck, -1, next.String())
	}
}
