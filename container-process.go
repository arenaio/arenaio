package main

import (
	"os/exec"
	"io"
)

type ContainerProcess struct {
	cmd *exec.Cmd
	stdIn io.WriteCloser
	stdOut io.ReadCloser
	stdErr io.ReadCloser
}

func NewContainerProcess(image string) (p *ContainerProcess, err error) {
	p = &ContainerProcess{
		cmd: exec.Command("docker", "run", "-ti", "--rm", image),
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

	return p, nil
}

func (p *ContainerProcess) Run() {

}

func (p *ContainerProcess) Kill() error {
	return p.cmd.Process.Kill()
}
