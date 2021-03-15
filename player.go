package main

import (
	"fmt"
	"math/rand"
	"spoof-game/util"
)

/**
 * Structure to represent a Player.
 * @attribute	initiator	True if the Player is starting the current round.
 * @attribute	id			The id of the Player.
 * @attribute	position	The position of the Player in the current round.
 * @attribute	coins		The coins picket by the Player in the current round.
 * @attribute 	guess		The guessed total coins on the table in the current round.
 */
type Player struct {
	initiator                  bool
	id, position, coins, guess int
}

/**
 * Guess the total amount of coins on the table in this round.
 * @param guesses			The array of guesses ordered by position of the players.
 * @param numberOfPlayers	The number of players currently playing.
 * @return 					The number of coins guessed by the player.
 */
func (player Player) guessCoins(guesses []int, numberOfPlayers int) int {
	maxTable := maxCoins*(numberOfPlayers-1) + player.coins
	minTable := minCoins*(numberOfPlayers-1) + player.coins
	for {
		index := rand.Intn(maxTable-minTable) + minTable
		if !util.IsInList(guesses[:player.position], index) { // up to position, to avoid 0s of the initial array
			return index
		} else {
			fmt.Printf("Player %d: %d is already in the list\n", player.position, index)
		}
	}
}

/**
 * Print in output a string declaring the Player name and position.
 * @param string	The desired output string.
 */
func (player Player) talk(text string) {
	outString := fmt.Sprintf("Player %d (pos %d): %s", player.id, player.position, text)
	if player.initiator {
		outString += " <- initiator"
	}
	outString += "\n"
	fmt.Printf(outString)
}

/**
 * Pick a random number of coins from maxCoins to minCoins.
 * @return	The picked number of coins.
 */
func drawCoins() int {
	return rand.Intn(maxCoins+1-minCoins) + minCoins
}

/**
 * Find the winner of the round.
 * @param guesses			The array of guesses ordered by position of the players.
 * @param numberOfPlayers	The number of players currently playing.
 * @return					The position of the winner of the round.
 */
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
