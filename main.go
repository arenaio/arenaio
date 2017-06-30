package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	//"bufio"
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
	fmt.Printf("Referee:  %s\nPlayer 1: %s\nPlayer 2: %s\n", *referee, *player1, *player2)

	p1, err := NewContainerProcess(*player1)

	err = p1.cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}

	//go func() {
	//	//defer p1stdin.Close()
	//	line := fmt.Sprintf("%d", 1)
	//	fmt.Printf("StdIn: %s", line)
	//	//io.WriteString(p1stdin, line)
	//	writer := bufio.NewWriter(p1stdin)
	//	writer.WriteString(line)
	//}()
	//
	//go func() {
	//	scanner := bufio.NewScanner(p1stdout)
	//	for scanner.Scan() {
	//		fmt.Printf("StdOut: %s\n", scanner.Text())
	//	}
	//	if err := scanner.Err(); err != nil {
	//		fmt.Fprintln(os.Stderr, "reading standard output:", err)
	//	}
	//}()
	//
	//go func() {
	//	scanner := bufio.NewScanner(p1stderr)
	//	for scanner.Scan() {
	//		fmt.Printf("StdErr: %s\n", scanner.Text())
	//	}
	//	if err := scanner.Err(); err != nil {
	//		fmt.Fprintln(os.Stderr, "reading standard error:", err)
	//	}
	//}()
	//
	//err = p1.Wait()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
