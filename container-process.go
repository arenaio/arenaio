package main

import (
	"github.com/arenaio/arenaio/game/bot"
)

func NewContainerProcess(image string, idx int) (p *bot.Process, err error) {
	return bot.NewProcess("docker", idx, "run", "-m 512m", "-c 1", "-i", "--rm", image)
}
