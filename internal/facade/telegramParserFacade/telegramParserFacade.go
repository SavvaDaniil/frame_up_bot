package telegramParserFacade

import (
	"fmt"
	"frame_up_bot/internal/middleware/telegramParserMiddleware"
	"frame_up_bot/internal/model/telegramParserModel"
	"frame_up_bot/internal/repository/nominationRepository"
	"frame_up_bot/internal/repository/registrationRepository"
	"frame_up_bot/internal/service/telegramParserService"
	"log"
	"strconv"
	"time"
)

func StartAsyncListener(telegramParser *telegramParserModel.TelegramParser, isOffsetAvailable bool) {
	var trying int = 0

	for {
		getUpdatesModel, err := telegramParserService.RequestGetUpdates(telegramParser, isOffsetAvailable)
		if err != nil && trying > 5 {
			log.Printf("Провальная попытка получить обновление с телеграма: %s", err.Error())
			panic("Провальная попытка получить обновление с телеграма")
		} else if err != nil {
			trying += 1
			log.Printf("Произошла ошибка соединения, ждем 10 сек и пробуем снова: %s", err.Error())
			fmt.Println("Произошла ошибка соединения, ждем 10 сек и пробуем снова")
			time.Sleep(10 * time.Second)
			continue
		}

		if getUpdatesModel.Result != nil {
			for _, update := range getUpdatesModel.Result {
				chat_id, isAccessGranted, err := telegramParserMiddleware.GetUserIdFromMessageAndCheckForAccess(&update.Message)
				if err != nil {
					log.Printf("Что-то случилось с проверкой id пользователя с массивом в кофигурации: %s", err.Error())
					panic("Что-то случилось с проверкой id пользователя с массивом в кофигурации")
				}
				if isOffsetAvailable {
					telegramParser.Offset = update.UpdateId + 1
				}
				if !isAccessGranted || *chat_id == 0 {
					continue
				}

				if update.Message.Text == "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" {

					statisticsAnswer, err := getStatisticsAboutRegistrationMessage()
					if err != nil {
						panic("Ошибка при попытке собрать статистику")
					}

					//fmt.Println(*statisticsAnswer)

					errSending := telegramParserService.SendHTMLTextMessage(telegramParser, *chat_id, *statisticsAnswer)
					if errSending != nil {
						log.Printf("Ошибка при попытке отправить сообщения: %s", err.Error())
						panic("Ошибка при попытке отправить сообщения")
					}

					/*
						errSending := telegramParserService.SendHTMLTextMessage(telegramParser, *chat_id, "Успешно")
						if errSending != nil {
							log.Printf("Ошибка при попытке отправить сообщения: %s", err.Error())
							panic("Ошибка при попытке отправить сообщения")
						}
					*/
				}

			}
		}

		trying = 0
		time.Sleep(1 * time.Second)

	}

}

func getStatisticsAboutRegistrationMessage() (*string, error) {

	nominations, err := nominationRepository.ListAllActive()
	if err != nil {
		log.Printf("Ошибка при попытке выборки номинаций: %s", err.Error())
		fmt.Printf("Ошибка при попытке выборки номинаций: %s", err.Error())
		return nil, err
	}

	registrations, err := registrationRepository.ListAllAfter2023()
	if err != nil {
		log.Printf("Ошибка при попытке выборки регистраций: %s", err.Error())
		fmt.Printf("Ошибка при попытке выборки номинаций: %s", err.Error())
		return nil, err
	}

	date_now := time.Now().Format("2006-01-02 15:04:05")
	var textAnswer string = "На дату: " + date_now + "\r\n"
	var registrationCount int = 0
	for _, nomination := range *nominations {
		registrationCount = 0
		for _, registration := range *registrations {
			if registration.NominationId == nomination.Id {
				registrationCount += 1
			}
		}
		textAnswer += "Номинация '" + *nomination.Name + "': " + strconv.Itoa(registrationCount) + "\r\n"
	}

	return &textAnswer, nil
}
