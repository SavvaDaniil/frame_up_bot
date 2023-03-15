package config

//"github.com/spf13/viper"

const DbConnection string = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

var GLOBAL_configOfApp Config

//var GLOBAL_botOptions BotOptions

type TG struct {
	Token string
}
type Auth struct {
	SecretKey string
}

type Config struct {
	TG            TG
	Auth          Auth
	UserWhiteList []int64
	URL           string
}

func ParseConfig() (Config, error) {

	TG := TG{
		Token: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	}
	Auth := Auth{
		SecretKey: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	}
	var userWhiteList []int64 = make([]int64, 2)
	userWhiteList = append(userWhiteList, 251535022) // Савва
	userWhiteList = append(userWhiteList, 262408249) //Белова

	var c Config = Config{
		TG:            TG,
		Auth:          Auth,
		UserWhiteList: userWhiteList,
		URL:           "",
	}
	return c, nil
}
