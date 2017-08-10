package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/arenaio/arenaio/pkg/bot"
)

// arenaio -r=tictactoe -p1=theNewB:12 -p2=KB:7

func main() {
	referee := flag.String("r", "", "a string")
	player1 := flag.String("p1", "", "a string")
	player2 := flag.String("p2", "", "a string")

	flag.Parse()

	if len(*referee) == 0 {
		log.Fatalln("Referee is required")
		os.Exit(1)
	}

	if len(*player1) == 0 {
		log.Fatalln("Player 1 is required")
		os.Exit(1)
	}

	if len(*player2) == 0 {
		log.Fatalln("Player 2 is required")
	}

	run(referee, player1, player2)
}

func run(referee, player1, player2 *string) {
	//fmt.Printf("Referee:  %s\nPlayer 1: %s\nPlayer 2: %s\n", *referee, *player1, *player2)

	p1, err := bot.NewContainerProcess(*player1, 1)
	if err != nil {
		log.Fatalln(err)
	}

	// init Referee

	// TODO: Ask Referee what to do next
	// TODO: switch
	// case EndOfGame
	// case pX send input
	// case pX request output
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for p1.Alive() {
			_, err := p1.Write([]byte(fmt.Sprintln("1")))
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	go func() {
		for p1.Alive() {
			var line string
			_, err := fmt.Fscan(p1, &line)
			if err != nil {
				log.Println("Err:", err)
			}
			fmt.Print(line)
		}
		wg.Done()
	}()
	wg.Wait()
}
