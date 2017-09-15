package telbot

import (
	"github.com/yanzay/tbot"
	"strconv"
	"time"
	"fmt"
)

func TimerHandler(m *tbot.Message) {
	// m.Vars contains all variables, parsed during routing
	secondsStr := m.Vars["seconds"]
	// Convert string variable to integer seconds value
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		m.Reply("Invalid number of seconds")
		return
	}
	m.Replyf("Timer for %d seconds started", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	m.Reply("Time out!")
}

func PhotoHandler(message *tbot.Message) {
	message.ReplyPhoto("./telbot/pics/photo.jpeg", "it's me")
}

func KeyboardHandler(message *tbot.Message) {
	buttons := [][]string{
		{"This", "Is", "A"},
		{"Way"},
	}
	message.ReplyKeyboard("Buttons example", buttons)
}

func HiHandler(message *tbot.Message) {
	// Handler can reply with several messages
	message.Replyf("Hello, %s!", message.From)
	userName := message.From
	time.Sleep(1 * time.Second)
	fmt.Println(userName)
	message.Reply("What's up?")
}

func EchoHandler(message *tbot.Message) {
	message.Reply(message.Text())
}

