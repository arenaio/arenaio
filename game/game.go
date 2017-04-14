package game

import (
	"math/rand"
	"time"
)

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
	GetMillisTimeForFirstRound() int
	InitReferee(playerCount int, prop Properties)
	GetFrameDataForView(round int, frame int, keyFrame bool) []string
	GetInitDataForView() []string
}

func Run(g Game, botBinaries []string) int {
	botCount := len(botBinaries)

	// shuffle bot positions
	botIndexes := make([]int, botCount)
	for i := 0; i < botCount; i++ {
		botIndexes[i] = i
	}
	for i := botCount - 1; i > 0; i-- {
		j := rand.Int() % (i + 1)
		botBinaries[i], botBinaries[j] = botBinaries[j], botBinaries[i]
		botIndexes[i], botIndexes[j] = botIndexes[j], botIndexes[i]
	}

	// TODO: initialize bot processes
	// TODO: init pipes

	winner := -1
	for round := 1; round <= g.GetMaxRoundCount(botCount); round++ {
		firstRound := round == 1

		for b := 0; b < botCount; b++ {
			if firstRound {
				// TODO: start process
				g.GetInitInputForPlayer(0)
			}

			g.GetInputForPlayer(round, 0)

			// TODO: send input
			// TODO: wait for output
			// TODO: handle timeout ?
		}

		// TODO: update gme & check end state
		g.UpdateGame(round)
		time.Sleep(time.Millisecond)
	}

	// TODO: evaluate winner
	winner = rand.Int()%(len(botBinaries)+1) - 1

	// no mapping required for draw games
	if winner == -1 {
		return winner
	}

	// map winner index id to bot index
	return botIndexes[winner]
}
