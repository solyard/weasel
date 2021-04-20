package main

import (
	"github.com/weasel/pkg/api"
	"github.com/weasel/pkg/telegram"
)

func main() {
	go api.InitialiseAPI()
	telegram.StartTelegramBot()

}
