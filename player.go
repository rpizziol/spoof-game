package main

import (
	"fmt"
	"math/rand"
	"spoof-game/util"
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

func (player Player) guessCoins(guesses []int, numberOfPlayers int) int {
	maxTable := maxCoins*(numberOfPlayers-1) + player.coins
	minTable := minCoins*(numberOfPlayers-1) + player.coins
	for {
		index := rand.Intn(maxTable-minTable) + minTable
		if !isInList(guesses[:player.position], index) { // up to position, to avoid 0s of the initial array
			return index
		} else {
			fmt.Printf("Player %d: %d is already in the list\n", player.position, index)
		}
	}
}

func (player Player) talk(text string) {
	outString := fmt.Sprintf("Player %d (pos %d): %s", player.id, player.position, text)
	if player.initiator {
		outString += " <- initiator"
	}
	outString += "\n"
	fmt.Printf(outString)
}

func (player Player) step(step int, numberOfPlayers int) {
	player.position = (player.position + numberOfPlayers + step) % numberOfPlayers
}

func drawCoins() int {
	return rand.Intn(maxCoins+1-minCoins) + minCoins
}

func (player Player) printPlayer() {
	player.talk(fmt.Sprintf("picked %d coins", player.coins))
}

func (player Player) findWinner(guesses []int, numberOfPlayers int) int {
	player.talk(fmt.Sprintf("guesses = %v", guesses))
	var distance = make([]int, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		distance[i] = util.Abs(guesses[len(guesses)-1] - guesses[i])
	}
	player.talk(fmt.Sprintf("distance = %v", distance))
	var winner = 0
	for i := 0; i < numberOfPlayers; i++ {
		if distance[i] < distance[winner] {
			winner = i
		}
	}
	return winner
}
