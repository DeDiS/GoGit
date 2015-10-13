package main

import (
	"math/rand"
	"fmt"
	"bufio"
	"os"
	"strings"
)

const (
	ROCK = iota
	SCISSORS
	PAPER
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Chose either (r)ock, (p)aper or (s)cissors: ")
		input, _ := reader.ReadString('\n')
		if human := Convert(string(input[0])); human >= 0{
			computer := rand.Intn(3)
			switch Wins(human, computer){
			case -1:
				fmt.Println("Sorry, the computer wins")
			case 0:
				fmt.Println("Draw, try again")
			case 1:
				fmt.Println("YES! You win")
			}
		}
	}
}

// Takes a string and converts it to one of ROCK, PAPER or SCISSORS
// Valid strings are both capital and lowercase:
// * r - ROCK
// * s - SCISSORS
// * p - PAPER
// Return -1 for non-valid string
func Convert(s string) int{

	s2 := strings.ToLower(s)

	switch s2 {
		case "r":
			return ROCK
		case "p":
			return PAPER
		case "s":
			return SCISSORS
	}

	return -1
}

// Takes two turns and returns
// * -1 for win of player2
// * 0 for draw
// * 1 for win of player 1
// Only valid turns are given in the argument
func Wins(player1, player2 int) int{

	switch player1 {
		case ROCK:
			switch player2 {
				case PAPER:
					return -1
				case SCISSORS:
					return 1
			}
		case PAPER:
			switch player2 {
				case ROCK:
					return 1
				case SCISSORS:
					return -1
			}
		case SCISSORS:
			switch player2 {
				case ROCK:
					return -1
				case PAPER:
					return 1
			}
	}

	return 0
}
