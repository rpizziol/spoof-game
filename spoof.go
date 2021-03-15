package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Spoof Game Problem
Spoof is a game played to establish who buys the next round in a pub (the looser pays for all).
Each player draws some number of coins between zero and three and holds them concealed. The game
proceeds in rounds: in a round each player guesses the total number of coins contained in all the
players’ hands. A player acts as the initiator and starts the guessing process that proceeds
clockwise until all players have guessed a number that hasn’t previously been taken. After the
guesses all players reveal their coins, the total is calculated and the player who choose the
closest number wins the round and exists the game. The others will play more rounds until only
one player is left. This player is who buys the drinks. Every time a round is played the first
player to guess a number rotates clockwise to the next available player.*/

const totalNumberOfPlayers = 10
const maxCoins = 3 // Maximum number of coins per player (default: 3)
const minCoins = 0 // Minimum number of coins per player (default: 0)

/**
 *
 * @param routineId
 * @param routinePosition
 * @param initiatorPosition
 * @param round
 * @param inputChannel
 * @param outputChannel
 * @param winnerChannel
 * @return
 */
func routineJob(routineId int, routinePosition int, initiatorPosition int, round int, inputChannel <-chan []int,
	outputChannel chan<- []int, winnerChannel chan<- int) {
	numberOfPlayers := totalNumberOfPlayers - round
	player := Player{id: routineId, position: routinePosition, coins: drawCoins()}
	player.initiator = player.position == initiatorPosition
	player.talk(fmt.Sprintf("picked %d coins", player.coins))
	var winner int
	if player.initiator {
		// The element passed between players (array of guesses + money box)
		// (NB The last place is reserved to the real value, updated by each player)
		var guesses = make([]int, numberOfPlayers+1)
		player.guess = player.guessCoins(guesses, numberOfPlayers)
		guesses[player.position] = player.guess
		// Update the overall value of coins on the table
		guesses[len(guesses)-1] = player.coins
		// Pass array of guesses to the next player
		outputChannel <- guesses
		// Wait for the guesses list to be updated by all players
		guesses = <-inputChannel
		player.talk(fmt.Sprintf("%v received", guesses))
		// Guesses round is over
		winner = player.findWinner(guesses, numberOfPlayers)
		player.talk(fmt.Sprintf("the winner of round %d is Player in position %d", round, winner))
		winnerChannel <- winner
	} else {
		guesses := <-inputChannel
		player.talk(fmt.Sprintf("%v received", guesses))
		// Send guesses array on the sending channel
		player.guess = player.guessCoins(guesses, numberOfPlayers)
		guesses[player.position] = player.guess
		// Update the overall value of coins on the table
		guesses[len(guesses)-1] += player.coins
		outputChannel <- guesses
	}
}

func main() {
	// Initialization procedures
	rand.Seed(time.Now().UnixNano())

	// Create list of player ids
	var playerIds = []int{0}
	// Add all remaining players to the list of ids
	for i := 1; i < totalNumberOfPlayers; i++ {
		playerIds = append(playerIds, i)
	}

	initiatorPosition := 0
	for round := 0; round < totalNumberOfPlayers-1; round++ {
		fmt.Printf("############ MASTER: ROUND %d\nPlayers in game: %v\n", round, playerIds)
		numberOfPlayers := totalNumberOfPlayers - round

		// Ring channels for communication between players
		var channel = make([]chan []int, numberOfPlayers)
		for i := range channel {
			channel[i] = make(chan []int)
		}

		// Channel to communicate winner of a round to the master
		var winnerChannel = make(chan int)

		// Move the initiator to the first position in case of need
		if initiatorPosition >= numberOfPlayers {
			initiatorPosition = 0
		}

		// Execute a routine for each player still in game
		for i := 0; i < numberOfPlayers; i++ {
			go routineJob(playerIds[i], i, initiatorPosition, round, channel[((i-1)+numberOfPlayers)%numberOfPlayers],
				channel[i], winnerChannel)
		}
		// Receive winner from the initiator
		winner := <-winnerChannel
		fmt.Printf("############ MASTER: winner is player in position %d\n\n", winner)

		for i := 0; ; i++ {
			if i == winner {
				// Remove winner player from table
				copy(playerIds[i:], playerIds[i+1:])
				playerIds = playerIds[:len(playerIds)-1]
				break
			}
		}
		closeChannels(channel)
		initiatorPosition++
	}
	fmt.Printf("############ MASTER: Player %d has to pay a drink for all!\n", playerIds[0])
}

/**
 * Close all channels in an array of channels.
 * @param channel	The array of channels to close.
 */
func closeChannels(channel []chan []int) {
	for c := range channel {
		close(channel[c])
	}
}
