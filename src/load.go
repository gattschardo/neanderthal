package main

import (
	"os"
	"bufio"
	"fmt"
)

func readDeckFile(filename string) (d *Deck) {

	d = new(Deck)
	last := 0
	col := Color(999)

	fh, ok := os.Open(filename, os.O_RDONLY, 0)

	if ok != nil {
		fmt.Fprintf(os.Stderr, "Could not open deck file `%s' for reading\n", filename)
		os.Exit(1)
	}

	f := bufio.NewReader(fh)

	for line, ok := f.ReadString('\n'); ok == nil; line, ok = f.ReadString('\n') {
		if last > deckSize {
			fmt.Fprintf(os.Stderr, "Warning: stopping after reading %d cards\n", deckSize)
			break
		}

		k := len(line) - 2
		if line[0] == '[' && line[k] == ']' {
			for i, s := range colorName {
				if s == line[1:k] {
					col = Color(i)
				}
			}
		} else if line != "\n" && line[0] != '#' {
			d[last] = Card{line[0 : k+1], col, 0}
			last++
		}
	}

	if ok == os.EOF {
		fmt.Fprintln(os.Stderr, "EOF encountered, returning")
	} else if ok != nil {
		fmt.Fprintf(os.Stderr, "Error while reading from deck file `%s': %s\n",
			filename, ok.String())
		os.Exit(1)
	}

	if last != deckSize {
		fmt.Fprintf(os.Stderr, "Warning: read %d cards instead of %d\n", last, deckSize)
	}

	return d
}
