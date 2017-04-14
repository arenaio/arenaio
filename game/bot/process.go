package bot

import (
	"io"
	"os"
	"os/exec"
	"sync"
)

type Process struct {
	m            sync.Mutex
	name         string
	processState *os.ProcessState
	inPipe       io.WriteCloser
	outPipe      io.Reader
}

// Creates a new bot process with given binary name
func NewProcess(name string, args ...string) (*Process, error) {
	cmd := exec.Command(name, args...)
	pr, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	pw, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// TODO create a per Bot error chan for colorization later on
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// wire up the process with a pipe reader/writer
	process := &Process{
		name:    name,
		inPipe:  pr,
		outPipe: pw,
	}

	// non-blocking call to wait
	go process.wait(cmd)

	return process, nil
}

// create a blocking call to process.wait to get
// notified via Alive method if the process exited somehow
func (p *Process) wait(cmd *exec.Cmd) {
	state, err := cmd.Process.Wait()
	if err != nil {
		println(err)
		return
	}
	p.m.Lock()
	p.processState = state
	p.m.Unlock()
}

func (p *Process) Alive() bool {
	p.m.Lock()
	defer p.m.Unlock()
	if p.processState != nil {
		return !p.processState.Exited()
	}
	return true
}

// io.Reader interface
// Reads the bot output
func (p *Process) Read(b []byte) (int, error) {
	return p.outPipe.Read(b)
}

// io.Writer interface
// Sends bytes to bot input
func (p *Process) Write(b []byte) (int, error) {
	return p.inPipe.Write(b)
}

// io.Closer interface
func (p *Process) Close() error {
	return p.inPipe.Close()
}
