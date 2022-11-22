package main

import (
	"fmt"

	"github.com/xtsfunral/discord-ping/bot"
	"github.com/xtsfunral/discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan chan struct{})
	return
}
