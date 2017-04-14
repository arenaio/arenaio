package main

import (
	"fmt"
	"os"

	"github.com/KevinBusse/arenaio/game"
	"github.com/KevinBusse/arenaio/game/tictactoe"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Program requires at least game name\n\n\tUsage: code-runner [Game] [...Bot]")
		os.Exit(1)
	}

	var g game.Game

	switch os.Args[1] {
	case "tictactoe":
		g = tictactoe.New()
	default:
		fmt.Fprintf(os.Stderr, "Game %q is unknown.\n", os.Args[1])
		os.Exit(2)
	}

	if len(os.Args) < 2+g.GetMinimumPlayerCount() {
		fmt.Fprintf(os.Stderr, "The game %s requires at least %d players.\n", g.GetGameName(), g.GetMinimumPlayerCount())
		os.Exit(3)
	}

	botBinaries := os.Args[2:]

	// Check that AI binaries are present
	for _, bot := range botBinaries {
		if _, err := os.Stat(bot); os.IsNotExist(err) {
			fmt.Printf("Bot %q couldn't be found.\n", bot)
			os.Exit(4)
		}
	}

	botScores := make([]float64, len(botBinaries)+1)

	fmt.Printf("%q", botBinaries[0])
	for i := 1; i < len(botBinaries); i++ {
		fmt.Printf(" | %q", botBinaries[i])
	}
	fmt.Println(" | Draw")

	games := 0.0
	for {
		games++
		result, err := game.Run(g, botBinaries)
		if err != nil {
			println(err)
			os.Exit(5)
		}

		if result == -1 {
			botScores[len(botScores)-1]++
		} else {
			botScores[result]++
		}

		fmt.Printf("%.1f%%", 100*botScores[0]/games)
		for i := 1; i < len(botScores); i++ {
			fmt.Printf(" | %.1f%%", 100*botScores[i]/games)
		}
		fmt.Println()
	}
}
