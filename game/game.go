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
	InitReferee(playerCount int, prop Properties)
	GetFrameDataForView(round int, frame int, keyFrame bool) []string
	GetInitDataForView() []string
}

func Run(g Game, botBinaries []string) int {
	// shuffle bot positions
	botIndexes := make([]int, len(botBinaries))
	for i := 0; i < len(botIndexes); i++ {
		botIndexes[i] = i
	}
	for i := len(botIndexes) - 1; i > 0; i-- {
		j := rand.Int() % (i + 1)
		botBinaries[i], botBinaries[j] = botBinaries[j], botBinaries[i]
		botIndexes[i], botIndexes[j] = botIndexes[j], botIndexes[i]
	}

	time.Sleep(100 * time.Millisecond)
	winner := rand.Int()%(len(botBinaries)+1) - 1

	// no mapping required for draw games
	if winner == -1 {
		return winner
	}

	// map winner index id to bot index
	return botIndexes[winner]
}
