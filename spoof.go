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
const minCoins = 0

// TODO close channels

func routineJob(routineId int, routinePosition int, inputChannel <-chan []int, outputChannel chan<- []int, round int, winnerChannel chan<- int, initiatorPosition int) {
	numberOfPlayers := totalNumberOfPlayers - round
	player := Player{id: routineId, position: routinePosition, coins: drawCoins()}
	player.initiator = player.position == initiatorPosition //round%numberOfPlayers
	player.printPlayer()
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

/*func routineJob(routineId int, channel []chan []int) {
	// Initialize
	player := Player{id: routineId, position: routineId}
	player.printPlayer()

	for round := 0; round < totalNumberOfPlayers-1; round++ {
		var winner int
		player.coins = drawCoins()
		numberOfPlayers := totalNumberOfPlayers - round
		player.talk(fmt.Sprintf("round %d", round))
		player.initiator = player.position == (round)
		if player.initiator {
			player.talk(fmt.Sprintf("Initiator of round %d", round))
		}
		player.printPlayer()
		if player.initiator {
			// The element passed between players (array of guesses + money box)
			// (NB The last place is reserved to the real value, updated by each player)
			var guesses = make([]int, numberOfPlayers+1)
			player.guess = player.guessCoins(guesses, numberOfPlayers)
			guesses[player.position] = player.guess
			// Update the overall value of coins on the table
			guesses[len(guesses)-1] = player.coins
			// Pass array of guesses to the next player
			channel[player.position] <- guesses
			guesses = <-channel[(player.position+numberOfPlayers-1)%numberOfPlayers]
			player.talk(fmt.Sprintf("%v received", guesses))
			// Guesses round is over
			winner = player.findWinner(guesses, numberOfPlayers)
			player.talk(fmt.Sprintf("the winner of round %d is Player in position %d", round, winner))
			// Send winner to all players
			winArray := []int{winner}
			for i := range channel[:numberOfPlayers] {
				if i != (player.position+numberOfPlayers-1)%numberOfPlayers && i != player.position {
					channel[i] <- winArray
				}
			}
			channel[player.position] <- winArray // Notify next initiator for last
		} else {
			// Wait for message on the receiving channel
			guesses := <-channel[(player.position+numberOfPlayers-1)%numberOfPlayers]
			player.talk(fmt.Sprintf("%v received", guesses))
			// Send guesses array on the sending channel
			player.guess = player.guessCoins(guesses, numberOfPlayers)
			guesses[player.position] = player.guess
			// Update the overall value of coins on the table
			guesses[len(guesses)-1] += player.coins
			channel[player.position] <- guesses
			winArray := <-channel[(player.position+numberOfPlayers-1)%numberOfPlayers]
			fmt.Printf("winArray = %v +++++++++++++++++++++++++++++++++\n", winArray)
			winner = winArray[0]
		}
		if winner == player.position {
			player.talk(fmt.Sprintf("I am the winner of round %d!", round))
			//close(channel[numberOfPlayers-1])
			return
		} else if winner < player.position {
			player.talk(fmt.Sprintf("The winner I received is %d. I have to step back!", winner))
			//if winner == 0 && player.position == numberOfPlayers-1 {
			//	player.step(1, numberOfPlayers) // Step forward
			//} else if winner != 0 {
			player.step(-1, numberOfPlayers-1) // Step back
			//}
			player.talk("I stepped back")
		} else {
			player.talk(fmt.Sprintf("The winner I received is %d", winner))
		}
		if round == totalNumberOfPlayers-2 {
			player.talk(fmt.Sprintf("I have to pay drinks for all! GAME OVER!!!"))
		}
	}
}*/

func main() {
	//fmt.Println("Initializing Spoof Game...")
	rand.Seed(time.Now().UnixNano())
	var playerIds = []int{0}

	// Add all remaining players to the list of ids
	for i := 1; i < totalNumberOfPlayers; i++ {
		playerIds = append(playerIds, i)
	}

	initiatorPosition := 0
	for round := 0; round < totalNumberOfPlayers-1; round++ {
		fmt.Printf("############ MASTER: ROUND %d\nPlayers in game: %v\n", round, playerIds)

		numberOfPlayers := totalNumberOfPlayers - round
		// Initialize ring channels
		var channel = make([]chan []int, numberOfPlayers)
		for i := range channel {
			channel[i] = make(chan []int)
		}
		var winnerChannel = make(chan int)

		if initiatorPosition >= numberOfPlayers {
			initiatorPosition = 0
		}

		fmt.Printf("++++++++++ initiatorPosition = %d, round = %d\n", initiatorPosition, round)
		// Execute a routine for each player still in game
		for i := 0; i < numberOfPlayers; i++ {
			go routineJob(playerIds[i], i, channel[((i-1)+numberOfPlayers)%numberOfPlayers], channel[i], round, winnerChannel, initiatorPosition)
		}
		winner := <-winnerChannel
		fmt.Printf("############ MASTER: winner is player in position %d\n\n", winner)

		j := 0
		for i := 0; ; i++ {
			if j == winner && playerIds[i] != -1 {
				// Remove winner player from table
				copy(playerIds[i:], playerIds[i+1:])
				playerIds[len(playerIds)-1] = 0
				playerIds = playerIds[:len(playerIds)-1]
				break
			} else if playerIds[i] != -1 {
				j++
			}
		}

		initiatorPosition++
	}
	fmt.Printf("############ MASTER: Player %d has to pay a drink for all!\n", playerIds[0])

	/*// Array of channels
	var channel = make([]chan []int, totalNumberOfPlayers)
	for i := range channel { // Initialize each single channel
		channel[i] = make(chan []int)
	}

	// Waitgroup for "barrier" syncronization
	var wg sync.WaitGroup

	// Initialize a routine for each player
	for i := 0; i < totalNumberOfPlayers; i++ {
		wg.Add(1)
		i := i
		go func() {
			routineJob(i, channel)
			wg.Done()
		}()
	}
	wg.Wait()

	//var round int
	//for round = 0; round < totalNumberOfPlayers-1; round++ {
	//	wg.Wait()
	//}*/
}
