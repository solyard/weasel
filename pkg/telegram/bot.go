package telegram

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/weasel/pkg/weasel"
	tb "gopkg.in/tucnak/telebot.v2"
)

var Bot *tb.Bot

func Start() {
	log.Info("Launching Telegram Bot ...")
	b, err := tb.NewBot(tb.Settings{
		Token:  weasel.LoadConfig(),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	Bot = b
	b.Start()
}

func SendAlert(alert string, channel string) {
	chatId, err := strconv.Atoi(channel)
	if err != nil {
		log.Errorf("Error while convert chat ID to INT: %v", err)
		return
	}
	Bot.Send(tb.ChatID(chatId), alert, tb.ParseMode("Markdown"))
}
