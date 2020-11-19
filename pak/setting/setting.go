package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"os"
)

//app节的配置
type App struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string

	DateStrFormat  string
	DateTimeFormat string
	TimeFormat     string
	DateFormat     string
}

var AppSetting = &App{}

type Serve struct {
	RunMode  string
	HttpPort int
}

var ServeSetting = &Serve{}

type DB struct {
	DB_IP   string
	DB_NAME string
}

var DBSetting = &DB{}

var Cfg *ini.File

func Setup() {
	var err error
	if len(os.Getenv("WAIHUI")) > 0 {
		fmt.Println("进入到/conf/conf.ini")
		Cfg, err = ini.Load("/conf/conf.ini")
	} else {
		fmt.Println("进入到/conf/confrelease.ini")
		Cfg, err = ini.Load("/conf/confrelease.ini")
	}
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/release.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServeSetting)
	mapTo("db", DBSetting)
}
func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
