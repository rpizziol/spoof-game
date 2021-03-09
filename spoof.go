package main

import (
	"fmt"
	"math/rand"
	"spoof-game/learning-go/util"
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

const numberOfPlayers = 5
const maxCoins = 4 // Maximum number of coins per player + 1 (default: 3)
const minCoins = 0

func routineJob(position int, channel []chan []int, guesses []int) {
	// Initialize
	player := Player{initiator: position == 0, position: position, coins: drawCoins()}
	fmt.Println(player.printPlayer())

	if player.initiator {
		player.guess = player.guessCoins(player.coins, guesses)
		guesses[position] = player.guess
		// Update the overall value of coins on the table
		guesses[len(guesses)-1] = player.coins
		// Pass array of guesses to the next player
		channel[position] <- guesses
		//close(channel[position])
		guesses := <-channel[(position+numberOfPlayers-1)%numberOfPlayers]
		fmt.Printf("Player %d: %v\n", player.position, guesses)
		// Guesses round is over
		var winner = findWinner(guesses)
		fmt.Printf("Player %d: the winner is Player %d\n", player.position, winner)
	} else {
		// Wait for message on the receiving channel
		guesses := <-channel[(position+numberOfPlayers-1)%numberOfPlayers]
		fmt.Printf("Player %d: %v received\n", player.position, guesses)
		// Send guesses array on the sending channel
		player.guess = player.guessCoins(player.coins, guesses)
		guesses[position] = player.guess
		// Update the overall value of coins on the table
		guesses[len(guesses)-1] += player.coins
		channel[position] <- guesses
	}
}

func findWinner(guesses []int) int {
	fmt.Printf("guesses: %v\n", guesses)
	var distance = make([]int, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		distance[i] = util.Abs(guesses[len(guesses)-1] - guesses[i])
	}
	fmt.Printf("distance: %v\n", distance)
	var winner = 0
	for i := 0; i < numberOfPlayers; i++ {
		if distance[i] < distance[winner] {
			winner = i
		}
	}
	return winner
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
