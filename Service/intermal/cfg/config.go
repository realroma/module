package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Cfg struct {
	Port   string
	DbName string
	DbHost string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("Serv")
	v.SetDefault("Port", "8080")
	v.SetDefault("DbName", "postgres")
	v.AutomaticEnv()

	var cfg Cfg

	return cfg
}

func (cfg *Cfg) GetDBSting() string {
	return fmt.Sprintf("postgres://%s@%s", cfg.DbHost, cfg.Port)
}
