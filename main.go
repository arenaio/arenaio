package main

import (
	"os"
	"fmt"
)

type Properties map[string]string

type Game interface {
	GetConfiguration() Properties
	GetInitInputForPlayer(playerIdx int) []string
	Prepare(round int)
	GetInputForPlayer(round int, playerIdx int)
	GetExpectedOutputLineCountForPlayer(playerIdx int) int
	HandlePlayerOutput(frame int, round int, playerIdx int, outputs []string)
	UpdateGame(round int)
	PopulateMessages(p Properties)
	GetGameName() string
	GetMinimumPlayerCount() int
	GetPlayerActions(playerIdx int, round int) []string
	IsPlayerDead(playerIdx int) bool
	GetDeathReason(playerIdx int) string
	GetScore(playerIdx int) int
	SetPlayerTimeout(frame int, round int, playerIdx int)
	GetMaxRoundCount(playerCount int) int
	GetMillisTimeForRound() int
	InitReferee(playerCount int, prop Properties)
	GetFrameDataForView(round int, frame int, keyFrame bool) [] string
	GetInitDataForView() []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Program requires at least game name")
		fmt.Fprintln(os.Stderr, "\n\tUsage: code-runner [Game] [...Bot]")
		os.Exit(1)
	}

	var game Game

	switch os.Args[1] {
	case "tictactoe":
		game = NewTicTacToe()
	default:
		fmt.Fprintf(os.Stderr, "Game %q is unknown.\n", os.Args[1])
		os.Exit(2)
	}

	if len(os.Args) < 2 + game.GetMinimumPlayerCount() {
		fmt.Fprintf(os.Stderr, "The game %s requires at least %d players.\n", game.GetGameName(), game.GetMinimumPlayerCount())
		os.Exit(3)
	}

	botBinaries := os.Args[2:]

	fmt.Printf("Testing AI %q", botBinaries[0])
	for i := 1; i < len(botBinaries); i++ {
		fmt.Printf(" vs %q", botBinaries[i])
	}
	fmt.Println()

	// Check that AI binaries are present
	for _, bot := range botBinaries {
		if _, err := os.Stat(bot); os.IsNotExist(err) {
			fmt.Printf("Bot %q couldn't be found.\n", bot)
			os.Exit(4)
		}
	}
}
