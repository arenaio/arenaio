package game

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/KevinBusse/arenaio/game/bot"
)

type Game interface {
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

func Run(g Game, botBinaries []string) (int, error) {
	botCount := len(botBinaries)

	// initialize bots
	bots := make(map[int]*bot.Process, botCount)
	for i := 0; i < botCount; i++ {
		bot, err := bot.NewProcess(botBinaries[i], i)
		if err != nil {
			return -1, err
		}
		bots[i] = bot
	}

	// shuffle bot positions
	for i := botCount - 1; i > 0; i-- {
		j := rand.Int() % (i + 1)
		bots[i], bots[j] = bots[j], bots[i]
	}

	winner := -1
	buffer := &bytes.Buffer{}
	for round := 1; round <= g.GetMaxRoundCount(botCount); round++ {
		firstRound := round == 1

		for b := 0; b < botCount; b++ {
			var lines []string
			if firstRound {
				lines = g.GetInitInputForPlayer(b)
			}
			lines = g.GetInputForPlayer(round, b)

			for _, line := range lines {
				buffer.WriteString(line)
			}

			n, err := buffer.WriteTo(bots[b])
			fmt.Fprintf(os.Stderr, "Debug: %d bytes written to bot %d\n", n, b)
			if err != nil {
				return -1, err
			}
			buffer.Reset()

			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				waitFor := 50 * time.Millisecond
				if firstRound {
					waitFor = 100 * time.Millisecond
				}
				time.Sleep(waitFor)
				wg.Done()
			}()
			wg.Wait()

			n, err = buffer.ReadFrom(bots[b])
			fmt.Fprintf(os.Stderr, "Debug: Read %d bytes from bot %d\n", n, b)
			if err != nil {
				return -1, err
			}

			lines = make([]string, 0)
			for {
				line, err := buffer.ReadString('\n')
				if err == io.EOF {
					break
				}
				lines = append(lines, line)
			}
			// describe frame?
			g.HandlePlayerOutput(0, round, b, lines)
			buffer.Reset()
		}

		// TODO: update gme & check end state
		g.UpdateGame(round)
	}

	// TODO: evaluate winner
	winner = rand.Int()%(len(botBinaries)+1) - 1

	// no mapping required for draw games
	if winner == -1 {
		return winner, nil
	}

	// map winner index id to bot index
	return bots[winner].Idx(), nil
}
