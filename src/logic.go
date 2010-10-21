package main

import "rand"

func setup() {
	rand.Seed(33)
}

func needed() Color {
	return Color(rand.Intn(4) + 2)
}
