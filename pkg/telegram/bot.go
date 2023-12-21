package telegram

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/weasel/pkg/weasel"
	tb "gopkg.in/telebot.v3"
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

	b.Handle("/chatID", func(c tb.Context) error {
		return c.Reply(fmt.Sprintf("Current Chat ID: %v", c.Chat().ID))
	})

	b.Handle("/threadID", func(c tb.Context) error {
		return c.Reply(fmt.Sprintf("Current Thread ID: %v \n\nIf you received 0 (zero) than you are in General topic or in chat without Threads", c.Message().ThreadID))
	})

	Bot = b
	b.Start()

}

func SendAlert(alert string, channel string) {
	chatId, err := strconv.Atoi(channel)
	if err != nil {
		log.Errorf("Error while convert chat ID to INT: %v", err)
		return
	}

	Bot.Send(tb.ChatID(chatId), alert, &tb.SendOptions{ParseMode: "Markdown"})
}

func SendAlertToTopic(alert string, channel string, topic string) {
	chatId, err := strconv.Atoi(channel)
	if err != nil {
		log.Errorf("Error while convert chat ID to INT: %v", err)
		return
	}

	topicID, err := strconv.Atoi(topic)
	if err != nil {
		log.Errorf("Error while convert topic ID to INT: %v", err)
	}

	Bot.Send(tb.ChatID(chatId), alert, &tb.SendOptions{ThreadID: topicID, ParseMode: "Markdown"})
}
