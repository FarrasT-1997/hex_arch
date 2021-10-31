package setting

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func GetConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 9000
	defaultConfig.Database.Driver = "mysql"
	defaultConfig.Database.Name = "buku"
	defaultConfig.Database.Address = "localhost"
	defaultConfig.Database.Port = 3306
	defaultConfig.Database.Username = "root"
	defaultConfig.Database.Password = "12345"

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./setting/")

	if err := viper.ReadInConfig(); err != nil {
		// log.Info("error to load config file, will use default value ", err)
		return &defaultConfig
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig
}
