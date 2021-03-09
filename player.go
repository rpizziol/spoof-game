package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Player struct {
	initiator                  bool
	id, position, coins, guess int
}

func isInList(list []int, elem int) bool {
	for _, a := range list {
		if a == elem {
			return true
		}
	}
	return false
}

func (player Player) guessCoins(myCoins int, guesses []int, position int) int {
	maxTable := maxCoins*(numberOfPlayers-1) + myCoins
	minTable := minCoins*(numberOfPlayers-1) + myCoins
	for {
		index := rand.Intn(maxTable-minTable) + minTable
		if !isInList(guesses[:position], index) { // up to position, to avoid 0s of the initial array
			return index
		} else {
			fmt.Printf("Player %d: %d is already in the list\n", player.position, index)
		}
	}
}

func drawCoins() int {
	return rand.Intn(maxCoins-minCoins) + minCoins
}

func (player Player) printPlayer() string {
	var output = "Player " + strconv.Itoa(player.id) + " (pos " + strconv.Itoa(player.position) + "): picked " +
		strconv.Itoa(player.coins) + " coins"
	if player.initiator {
		output += " <- initiator"
	}
	return output
}
