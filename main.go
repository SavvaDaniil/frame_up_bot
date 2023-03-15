package main

import (
	"fmt"
	"frame_up_bot/config"
	"frame_up_bot/internal/facade/telegramParserFacade"
	"frame_up_bot/internal/factory/telegramParserFactory"
	"log"
)

func main() {

	var errConfig error
	config.GLOBAL_configOfApp, errConfig = config.ParseConfig()
	if errConfig != nil {
		panic(fmt.Sprintf("Cannot parse config %s", errConfig.Error()))
	}

	telegramParser, errCreateTP := telegramParserFactory.Create(config.GLOBAL_configOfApp.TG.Token)
	if errCreateTP != nil {
		log.Printf("Fail create parser for telegram: %s", errCreateTP.Error())
		panic("Fail create parser for telegram")
	}

	telegramParserFacade.StartAsyncListener(telegramParser, true)
}
