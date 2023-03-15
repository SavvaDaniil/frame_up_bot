package telegramParserFactory

import (
	"fmt"
	"frame_up_bot/internal/model/telegramParserModel"
	"frame_up_bot/internal/service/telegramParserService"
	"log"
)

func Create(accessToken string) (*telegramParserModel.TelegramParser, error) {

	var telegramParser telegramParserModel.TelegramParser = telegramParserModel.TelegramParser{
		AccessToken: accessToken,
		Offset:      0,
	}

	var respGetMeModel telegramParserModel.GetMeModel = telegramParserModel.GetMeModel{}
	err := telegramParserService.RequestGetMe(&telegramParser, &respGetMeModel)
	if err != nil {
		panic("Не удалось инициализировать телеграм парсер")
	}
	if !respGetMeModel.IsOk || respGetMeModel.Result.Id == 0 {
		panic("Не удалось создать телеграм парсер, проверьте правильность AccessToken")
	}

	fmt.Println("Парсер телеграм создан")
	log.Printf("Парсер телеграм создан")

	return &telegramParser, nil
}
