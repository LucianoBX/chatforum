package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// 服务器信息
type App struct {
	Address  string
	Static   string
	Log      string
	Locale   string
	Language string
}

// 数据库配置
type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

// 整体配置
type Configuration struct {
	App          App
	Db           Database
	LocaleBundle *i18n.Bundle
}

var config *Configuration
var once sync.Once

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {

	// sync.Once 确保单例
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}

		// 本地化初始设置
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
		bundle.MustLoadMessageFile(config.App.Locale + "/active.en.json")
		bundle.MustLoadMessageFile(config.App.Locale + "/active." + config.App.Language + ".json")
		config.LocaleBundle = bundle
	})
	return config
}
