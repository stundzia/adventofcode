package day22

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strings"
)


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 22, "\n\n")
	player1 := utils.SliceStringToInt(strings.Split(input[0], "\n")[1:])
	player2 := utils.SliceStringToInt(strings.Split(input[1], "\n")[1:])
	game := NewGame(player1, player2, false)
	var gg bool
	var winner int
	for ;gg != true; {
		gg, winner = game.PlayRound()
	}
	score := game.getScore(winner)
	return fmt.Sprintf("%d", score)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 22, "\n\n")
	player1 := utils.SliceStringToInt(strings.Split(input[0], "\n")[1:])
	player2 := utils.SliceStringToInt(strings.Split(input[1], "\n")[1:])
	game := NewGame(player1, player2, true)
	var gg bool
	var winner int
	for ;gg != true; {
		gg, winner = game.PlayRound()
	}
	score := game.getScore(winner)
	return fmt.Sprintf("%d", score)
}
