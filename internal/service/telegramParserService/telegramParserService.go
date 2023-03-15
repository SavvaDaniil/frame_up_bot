package telegramParserService

import (
	"encoding/json"
	"errors"
	"fmt"
	"frame_up_bot/internal/model/telegramParserModel"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const telegramApiUrl string = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

func SendHTMLTextMessage(telegramParser *telegramParserModel.TelegramParser, chatId int, textMessage string) error {

	data := url.Values{
		"chat_id": {strconv.Itoa(chatId)},
		"text":    {textMessage},
		//"parse_mode": {"HTML"},
	}

	var url string = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	...
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("SendTextMessage: Провальная попытка прочитать json ответ: %s", err)
		return errors.New("SendTextMessage: Провальная попытка прочитать json ответ")
	}

	var getUpdatesModelResult *telegramParserModel.GetUpdatesModelResult
	if err := json.Unmarshal(body, &getUpdatesModelResult); err != nil { // Parse []byte to go struct pointer
		log.Printf("SendTextMessage: Can not unmarshal JSON: %s", err)
		fmt.Println("SendTextMessage: Can not unmarshal JSON")
		return err
	}
	fmt.Println("SendTextMessage: Сообщение успешно отправлено")

	return nil
}

func RequestGetUpdates(telegramParser *telegramParserModel.TelegramParser, isOffsetAvailable bool) (*telegramParserModel.GetUpdatesModel, error) {
	var queryOffset string = ""
	if isOffsetAvailable && telegramParser.Offset != 0 {
		queryOffset = "?offset=" + strconv.Itoa(telegramParser.Offset)
	}

	...

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("nil, GetUpdates: Провальная попытка прочитать json ответ")
	}

	var respGetUpdatesModel telegramParserModel.GetUpdatesModel
	if err := json.Unmarshal(body, &respGetUpdatesModel); err != nil {
		log.Printf("GetUpdates: Can not unmarshall JSON: %s", err.Error())
		return nil, errors.New("GetUpdates: Can not unmarshall JSON")
	}

	if respGetUpdatesModel.Result != nil {
		if len(respGetUpdatesModel.Result) > 0 {
			telegramParser.Offset = respGetUpdatesModel.Result[len(respGetUpdatesModel.Result)-1].UpdateId + 1
		}
	}

	return &respGetUpdatesModel, nil
}

func RequestGetMe(telegramParser *telegramParserModel.TelegramParser, respGetModel *telegramParserModel.GetMeModel) error {

	var url string = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("GetMe: провальная попытка установить соединение: %s", err.Error())
		return errors.New("GetMe: провальная попытка установить соединение")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("GetMe: Провальная попытка прочитать json ответ")
	}

	if err := json.Unmarshal(body, &respGetModel); err != nil {
		log.Printf("GetMe: Can not unmarshal JSON: %s", err.Error())
		fmt.Println("GetMe: Can not unmarshal JSON")
		return err
	}

	return nil
}
