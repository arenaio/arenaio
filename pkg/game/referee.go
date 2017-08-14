package game

type Referee interface {
	GetConfiguration() Properties
	GetInitInputForPlayer(playerIdx int) []string
	Prepare(round int)
	GetInputForPlayer(round int, playerIdx int) []string
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
