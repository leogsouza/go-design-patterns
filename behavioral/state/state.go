package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// GameState holds the GameContext
type GameState interface {
	executeState(*GameContext) bool
}

// GameContext represents the current context of a game
type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

// StartState represents the initial state of the game
type StartState struct{}

func (s *StartState) executeState(c *GameContext) bool {
	c.Next = &AskState{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)

	fmt.Println("Introduce a number of retries to set the difficulty:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	return true
}

// AskState represents the state where the game is asking for a number
type AskState struct {
}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n",
		c.Retries)

	var n int
	stdin := bufio.NewReader(os.Stdin)
	fmt.Fscanf(stdin, "%d", &n)
	c.Retries = c.Retries - 1

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}

	if c.Retries == 0 {
		c.Next = &FinishState{}
	}
	return true
}

// FinishState represents the final state of a game
type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		println("Congrats, you won")
	} else {
		println("You lose")
	}
	return false
}

func main() {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}
	for game.Next.executeState(&game) {
	}
}
