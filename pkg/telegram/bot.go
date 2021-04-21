package telegram

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/weasel/pkg/weasel"
	tb "gopkg.in/tucnak/telebot.v2"
)

var Bot *tb.Bot

func StartTelegramBot() {
	fmt.Println("Launching TelegramBot ...")
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		Token:  weasel.LoadConfig(),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	Bot = b
	b.Start()
}

func SendMessageToBot(alert string, channel string) {
	chatId, err := strconv.Atoi(channel)
	if err != nil {
		fmt.Printf("Error while convert chat ID to INT: %v", err)
	}
	Bot.Send(tb.ChatID(chatId), alert, tb.ParseMode("Markdown"))
}
