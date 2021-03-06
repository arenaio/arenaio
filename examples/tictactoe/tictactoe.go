package tictactoe

import (
	. "github.com/arenaio/arenaio/game"
	"github.com/arenaio/arenaio/pkg/game"
)

type TicTacToe struct{}

func New() *TicTacToe {
	return &TicTacToe{}
}

func (t *TicTacToe) GetConfiguration() game.Properties {
	return nil
}

func (t *TicTacToe) GetInitInputForPlayer(playerIdx int) []string {
	return nil
}

func (t *TicTacToe) Prepare(round int) {

}

func (t *TicTacToe) GetInputForPlayer(round int, playerIdx int) []string {
	return nil
}

func (t *TicTacToe) GetExpectedOutputLineCountForPlayer(playerIdx int) int {
	return -1
}

func (t *TicTacToe) HandlePlayerOutput(frame int, round int, playerIdx int, outputs []string) {

}

func (t *TicTacToe) UpdateGame(round int) {

}

func (t *TicTacToe) PopulateMessages(p game.Properties) {

}

func (t *TicTacToe) GetGameName() string {
	return "TicTacToe"
}

func (t *TicTacToe) GetMinimumPlayerCount() int {
	return 2
}

func (t *TicTacToe) GetPlayerActions(playerIdx int, round int) []string {
	return nil
}

func (t *TicTacToe) IsPlayerDead(playerIdx int) bool {
	return false
}

func (t *TicTacToe) GetDeathReason(playerIdx int) string {
	return ""
}

func (t *TicTacToe) GetScore(playerIdx int) int {
	return -1
}

func (t *TicTacToe) SetPlayerTimeout(frame int, round int, playerIdx int) {

}

func (t *TicTacToe) GetMaxRoundCount(playerCount int) int {
	return 9
}

func (t *TicTacToe) GetMillisTimeForRound() int {
	return 20
}

func (t *TicTacToe) GetMillisTimeForFirstRound() int {
	return 10
}

func (t *TicTacToe) InitReferee(playerCount int, prop game.Properties) {

}

func (t *TicTacToe) GetFrameDataForView(round int, frame int, keyFrame bool) []string {
	return nil
}

func (t *TicTacToe) GetInitDataForView() []string {
	return nil
}
