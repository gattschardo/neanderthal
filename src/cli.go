package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

const (
	columns = 8
	space = " "
	empty = "  "
	closedCard = "X"
	openCard = "[]"
	newline = "\n"
)

func printTable(d *Deck, open int, needed string) {
	opened := false

	for i, c := range(*d) {
		if c.owner > 0 {
			fmt.Print(empty)
		} else if i == open {
			fmt.Print(openCard)
			opened = true
		} else {
			fmt.Printf("%2d", i + 1)
		}

		if (i % columns) == columns - 1 {
			fmt.Print(newline)
		} else {
			fmt.Print(space)
		}
	}

	if opened {
		fmt.Println("Offene Karte:", d[open].String())
	}
	if needed != "" {
		fmt.Println("Gesucht:", needed)
	}
}

func promptCard() (c int) {
	c = -1

	in := bufio.NewReader(os.Stdin)

	for c < 0 {
		fmt.Print("Karte aufdecken> ")
		s, err := in.ReadString('\n')

		if s == "q\n" || err == os.EOF {
			fmt.Println()
			os.Exit(0)
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %s", err.String())
			os.Exit(1)
		}

		c, err = strconv.Atoi(s[0:len(s)-1])
		if err != nil || c < 1 || c > deckSize {
			fmt.Printf("Bitte gültige Zahl aus 1..%d eingeben!\n", deckSize)
			c = -1
		}
	}

	return c - 1
}
