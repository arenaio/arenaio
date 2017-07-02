package main

import (
	"os/exec"
	"io"
	//"fmt"
	"bufio"
	"log"
)

type ContainerProcess struct {
	cmd    *exec.Cmd
	stdIn  io.WriteCloser
	stdOut io.ReadCloser
	stdErr io.ReadCloser
	in     chan string
	out    chan string
	err    chan string
	done   chan struct{}
}

func handleWriteCloser(in <-chan string, w io.WriteCloser, doneChan <-chan struct{}) {
	done := false
	writer := bufio.NewWriter(w)

	for done != true {
		select {
		case line := <-in:
			log.Println("in:", line)
			//n, err := writer.WriteString(fmt.Sprintf("%s\n", line))
			n, err := writer.WriteString(line)
			log.Println(n, "bytes written", err)
		case <-doneChan:
			done = true
		}
	}
}

func handleReadCloser(r io.ReadCloser, out chan<- string, name string) {
	var err error
	reader := bufio.NewReader(r)
	line := ""

	for err == nil {
		buffer, isPrefix, err := reader.ReadLine()
		log.Println(len(buffer), "bytes read from", name, isPrefix, err)
		if err == nil {
			line += string(buffer)
			if !isPrefix {
				log.Println(name, ":", line)
				out <- line
				line = ""
			}
		}
	}
}

func NewContainerProcess(image string) (p *ContainerProcess, err error) {
	p = &ContainerProcess{
		cmd:  exec.Command("docker", "run", "-i", "--rm", image),
		in:   make(chan string),
		out:  make(chan string),
		err:  make(chan string),
		done: make(chan struct{}),
	}

	p.stdIn, err = p.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	p.stdOut, err = p.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	p.stdErr, err = p.cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = p.cmd.Start()
	if err != nil {
		return nil, err
	}

	go handleWriteCloser(p.in, p.stdIn, p.done)
	go handleReadCloser(p.stdOut, p.out, "out")
	go handleReadCloser(p.stdErr, p.err, "err")

	return p, nil
}

func (p *ContainerProcess) Kill() error {
	p.done <- struct{}{}
	p.stdIn.Close()
	p.stdOut.Close()
	p.stdErr.Close()
	return p.cmd.Process.Kill()
}
