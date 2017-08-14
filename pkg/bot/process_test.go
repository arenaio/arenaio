package bot_test

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/arenaio/arenaio/pkg/bot"
)

func TestNewBotProcess(t *testing.T) {
	bot, err := bot.NewProcess("sleep", 0, "1")
	if err != nil {
		t.Error(err)
	}

	if !bot.Alive() {
		t.Errorf("Expected to be alive.")
	}

	threshold := 10 * time.Millisecond
	time.Sleep(time.Second + threshold)

	if bot.Alive() {
		t.Errorf("Expected to be done.")
	}
}

func TestBotProcess_ReadWrite(t *testing.T) {
	bashBot, err := bot.NewProcess("bash", 0)
	if err != nil {
		t.Error(err)
	}

	// close bashBot after 10ms
	go func() {
		time.Sleep(10 * time.Millisecond)
		bashBot.Close()
	}()

	_, err = bashBot.Write([]byte("echo hello world"))
	if err != nil {
		t.Error(err)
	}

	bytes, err := ioutil.ReadAll(bashBot)
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != "hello world\n" {
		t.Errorf("Expected to get %q. Got %q.", "hello world", string(bytes))
	}
}
