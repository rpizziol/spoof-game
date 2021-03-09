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

const numberOfPlayers = 10
const maxCoins = 3 // Maximum number of coins per player (default: 3)
const minCoins = 0

// TODO close channels //close(channel[position])

func routineJob(position int, channel []chan []int, guesses []int) {
	// Initialize
	player := Player{id: position, initiator: position == 0, position: position, coins: drawCoins()}
	player.printPlayer()

	if player.initiator {
		player.guess = player.guessCoins(player.coins, guesses, player.position)
		guesses[position] = player.guess
		// Update the overall value of coins on the table
		guesses[len(guesses)-1] = player.coins
		// Pass array of guesses to the next player
		channel[position] <- guesses
		guesses := <-channel[(position+numberOfPlayers-1)%numberOfPlayers]
		player.talk(fmt.Sprintf("%v received", guesses))
		// Guesses round is over
		var winner = player.findWinner(guesses)
		if player.position == winner {
			player.talk("I am the winner!")
			// TODO exitRing()
		} else {
			player.talk(fmt.Sprintf("the winner is Player in position %d", winner))
			// Send winner to all players
			winArray := []int{winner}
			for i := range channel {
				channel[i] <- winArray
			}
		}
	} else {
		// Wait for message on the receiving channel
		guesses := <-channel[(position+numberOfPlayers-1)%numberOfPlayers]
		player.talk(fmt.Sprintf("%v received", guesses))
		// Send guesses array on the sending channel
		player.guess = player.guessCoins(player.coins, guesses, player.position)
		guesses[position] = player.guess
		// Update the overall value of coins on the table
		guesses[len(guesses)-1] += player.coins
		channel[position] <- guesses
		winArray := <-channel[(position+numberOfPlayers-1)%numberOfPlayers]
		winner := winArray[0]
		player.talk(fmt.Sprintf("%d received (winner)", winner))
	}
}

func main() {
	fmt.Println("Initializing Spoof Game...")
	rand.Seed(time.Now().UnixNano())
	// Array of channels
	var channel = make([]chan []int, numberOfPlayers)
	for i := range channel { // Initialize each single channel
		channel[i] = make(chan []int)
	}

	// The element passed between players (array of guesses + money box)
	// (NB The last place is reserved to the real value, updated by each player)
	var guesses = make([]int, numberOfPlayers+1)

	// Initialize a routine for each player
	for i := 0; i < numberOfPlayers; i++ {
		go routineJob(i, channel, guesses)
	}
	// TODO fix this with wait group?
	fmt.Scanln()
}
