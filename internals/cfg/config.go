package cfg

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct { //наша структура для хранения конфигов, я полагаюсь на Viper для матчинга имен
	Port        string
	DbName      string
	DbUser      string
	DbPass      string
	DBParseTime string
	DbHost      string
	DbPort      string
}

//ДОПИЛИТЬ

func LoadAndStoreConfig() Cfg {
	v := viper.New()             //создаем экземпляр нашего ридера для Env
	v.SetEnvPrefix("SERV")       //префикс, все переменные нашего сервера должны теперь стартовать с SERV_ для того, чтобы не смешиваться с системными
	v.SetDefault("PORT", "8080") //ставим умолчальные настройки
	v.SetDefault("DBUSER", "dclib_user")
	v.SetDefault("DBPASS", "password_")
	v.SetDefault("DBPARSETIME", "parseTime=True")
	v.SetDefault("DBHOST", "")
	v.SetDefault("DBPORT", "3306")
	v.SetDefault("DBNAME", "dclib_test")
	v.AutomaticEnv() //собираем наши переменные с системных

	var cfg Cfg

	err := v.Unmarshal(&cfg) //закидываем переменные в cfg после анмаршалинга
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

//pasha:password@/dclibrary?parseTime=true

func (cfg *Cfg) GetDBString() string { //маленький метод для сборки строки соединения с БД
	return fmt.Sprintf("%s:%s@/%s?%s", cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DBParseTime)
}
