package telegramParserModel

type TelegramParser struct {
	..............
}

type TelegramBotCommand struct {
	...............
}

type GetMeModelResult struct {
	Id    int  `json:"id"`
	IsBot bool `json:"is_bot"`
}

type GetMeModel struct {
	IsOk   bool             `json:"ok"`
	Result GetMeModelResult `json:"result"`
}

// MARK: GetUpdates
type GetUpdatesModelMessageFrom struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type GetUpdatesModelMessagePhoto struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type GetUpdatesModelMessage struct {
	MessageId int                           `json:"message_id"`
	UnixDate  int                           `json:"date"`
	Text      string                        `json:"text"`
	From      GetUpdatesModelMessageFrom    `json:"from"`
	Photo     []GetUpdatesModelMessagePhoto `json:"photo"`
	Caption   string                        `json:"caption"`
}

type GetUpdatesModelResult struct {
	UpdateId int                    `json:"update_id"`
	Message  GetUpdatesModelMessage `json:"message"`
}

type GetUpdatesModel struct {
	IsOk   bool                    `json:"ok"`
	Result []GetUpdatesModelResult `json:"result"`
}
