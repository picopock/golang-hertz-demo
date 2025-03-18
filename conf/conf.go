package conf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	Conf *Config
)

type Config struct {
	Env string

	App   App   `mapstructure:"app"`
	MySQL MySQL `mapstructure:"mysql"`
	Redis Redis `mapstructure:"redis"`
}

type MySQL struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	DB       string `mapstructure:"db"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type App struct {
	Port            uint16 `mapstructure:"port"`
	EnablePprof     bool   `mapstructure:"enable_pprof"`
	EnableGzip      bool   `mapstructure:"enable_gzip"`
	EnableAccessLog bool   `mapstructure:"enable_access_log"`
	LogLevel        string `mapstructure:"log_level"`
	LogFileName     string `mapstructure:"log_file_name"`
	LogMaxSize      int    `mapstructure:"log_max_size"`
	LogMaxBackups   int    `mapstructure:"log_max_backups"`
	LogMaxAge       int    `mapstructure:"log_max_age"`
}

func InitConf() {
	Conf = new(Config)
	err := viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
		os.Exit(1)
	}

	if Conf.App.Port == 0 {
		Conf.App.Port = 3000
	}
}
