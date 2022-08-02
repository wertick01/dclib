package cfg

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Cfg struct { //наша структура для хранения конфигов, я полагаюсь на Viper для матчинга имен
	Port   string `env:"PORT"`
	DbName string `env:"DBNAME"`
	DbUser string `env:"DBUSER"`
	DbPass string `env:"DBPASS"`
	//DBParseTime string `env-default:"parsetime=True"`
	DbHost string `env:"DBHOST"`
	DbPort string `env:"DBPORT"`
}

var configInstance *Cfg
var configErr error

func GetConfig() (*Cfg, error) {
	if configInstance == nil {
		var readConfigOnce sync.Once
		readConfigOnce.Do(func() {
			configInstance = &Cfg{}
			configErr = cleanenv.ReadEnv(configInstance)
		})
	}

	fmt.Printf("%+v\n", configInstance)

	return configInstance, configErr
}

func (cfg *Cfg) GetDBString() string { //маленький метод для сборки строки соединения с БД
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}
