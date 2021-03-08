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
const maxCoins = 4 // Maximum number of coins per player + 1 (default: 3)
const minCoins = 0

func threadJob(position int, channel []chan []int, guesses []int) {
	// Initialize
	myCoins := drawCoins()
	player := Player{initiator: position == 0, position: position, coins: myCoins}
	fmt.Println(player.printPlayer())

	if player.initiator {
		player.guess = player.guessCoins(myCoins, guesses)
		guesses[position] = player.guess
		// Pass array of guesses to the next player
		channel[position] <- guesses
		//close(channel[position])
		guesses := <-channel[(position+numberOfPlayers-1)%numberOfPlayers]
		//guesses = msg
		fmt.Printf("Player %d: %v\n", player.position, guesses)
	} else {
		// Wait for message on the receiving channel
		guesses := <-channel[(position+numberOfPlayers-1)%numberOfPlayers]
		//guesses = msg
		fmt.Printf("Player %d: %v\n", player.position, guesses)
		// Send guesses array on the sending channel
		player.guess = player.guessCoins(myCoins, guesses)
		guesses[position] = player.guess
		channel[position] <- guesses
	}

	//close(channel[position])
	// If initiator OR pushed
	// Guess number of coins in the table (a number not already guessed)
	// Push next thread

	// If initiator AND pushed
	// Declare guesses round is over

	// Add

	// Wait until all threads declared
	// Publish number of coins
	// Calculate total sum (and distance from my guess)

}

func main() {
	fmt.Println("Initializing Spoof Game...")
	rand.Seed(time.Now().UnixNano())
	// Initialize array of channels
	var channel = make([]chan []int, numberOfPlayers)
	// Initialize each single channel
	for i := range channel {
		channel[i] = make(chan []int)
	}
	var guesses = make([]int, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		go threadJob(i, channel, guesses)
	}
	fmt.Scanln()
	//fmt.Println(guesses)
}
