package main

import (
	"log"
	"os"
	"github.com/yanzay/tbot"
	"github.com/YAWAL/GoTraining/telbot"
)

func main() {
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/answer", "42")

	bot.Handle("yo", "YO!")
	bot.Handle("hi", "hi!")
	bot.HandleFunc("/hi", telbot.HiHandler)

	bot.HandleFunc("/timer {seconds}", telbot.TimerHandler)
	bot.HandleFunc("/photo", telbot.PhotoHandler)
	bot.HandleFunc("/keyboard", telbot.KeyboardHandler)

	// Set default handler if you want to process unmatched input
	bot.HandleDefault(telbot.EchoHandler)

	bot.ListenAndServe()

}
