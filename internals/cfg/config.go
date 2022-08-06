package cfg

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Cfg struct {
	Port   string `env:"PORT"`
	DbName string `env:"DBNAME"`
	DbUser string `env:"DBUSER"`
	DbPass string `env:"DBPASS"`
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

func (cfg *Cfg) GetDBString() string {
	//return "dclib_user:password_@tcp(localhost:3306)/dclib_test?parseTime=True"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}
