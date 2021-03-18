# Spoof Game Problem

Spoof is a game played to establish who buys the next round in a pub (the looser pays for all).

Each player draws some number of coins between zero and three and holds them concealed. The game proceeds in rounds: in a round each player guesses the total number of coins contained in all the players’ hands. A player acts as the initiator and starts the guessing process that proceeds clockwise until all players have guessed a number that hasn’t previously been taken. After the guesses all players reveal their coins, the total is calculated and the player who choose the closest number wins the round and exists the game. The others will play more rounds until only one player is left. This player is who buys the drinks. Every time a round is played the first player to guess a number rotates clockwise to the next available player.

## Requirements

It has been tested on Ubuntu 20.10, but should work on any OS with installed [Go](https://golang.org/).

## Installation

1. Get the code from [GitHub](https://github.com/rpizziol/spoof-game) or directly
clone it from the git repository:
```
git clone https://github.com/rpizziol/spoof-game.git
```

2. Eventually change the game variables in the header of the `spoof.go` file (`totalNumberOfPlayers`, `maxCoins` and `minCoins`).


## Usage
* Run the Go project (`go run .`)

## Authors

* **Roberto Pizziol** - *Spoof Game* - [rpizziol](https://github.com/rpizziol/)

