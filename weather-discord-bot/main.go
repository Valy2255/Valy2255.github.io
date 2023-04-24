package main

import (
	"fmt"
	
	"github.com/Valy2255/weather-bot/config"
	"github.com/Valy2255/weather-bot/bot"
	
)

func main() {
	err:=config.ReadConfig()

	if err!=nil {
		fmt.Println(err.Error())
	}
	bot.Start()
	
	<-make(chan struct{})
	
}