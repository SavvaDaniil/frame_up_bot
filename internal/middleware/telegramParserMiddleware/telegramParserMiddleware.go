package telegramParserMiddleware

import (
	"frame_up_bot/config"
	"frame_up_bot/internal/model/telegramParserModel"
)

func GetUserIdFromMessageAndCheckForAccess(getUpdatesModelMessage *telegramParserModel.GetUpdatesModelMessage) (*int, bool, error) {
	if (getUpdatesModelMessage.From == telegramParserModel.GetUpdatesModelMessageFrom{}) {
		var empty_user_id = 0
		return &empty_user_id, false, nil
	}

	for _, value := range config.GLOBAL_configOfApp.UserWhiteList {
		for int(value) == getUpdatesModelMessage.From.Id {
			return &getUpdatesModelMessage.From.Id, true, nil
		}
	}

	return &getUpdatesModelMessage.From.Id, false, nil
}
