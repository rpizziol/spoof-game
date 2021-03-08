package main

import "math/rand"

type Player struct {
	initiator              bool
	position, coins, guess int
}

func guessCoins(myCoins int, guesses []int) int {
	// TODO use 'guesses' to remove elements already in use
	maxTable := maxCoins*(numberOfPlayers-1) + myCoins
	minTable := minCoins*(numberOfPlayers-1) + myCoins
	return rand.Intn(maxTable-minTable) + minTable
}

func drawCoins() int {
	return rand.Intn(maxCoins-minCoins) + minCoins
}
