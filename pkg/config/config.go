package config

import (
	"github.com/BurntSushi/toml"
	"github.com/andy-smoker/wh-server/pkg/repository/postgres"
)

type ServerCFG struct {
	Addr string `toml:"addr"`
	Port string `toml:"port"`
}

type PostgresCFG struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

type Config struct {
	PGcfg  postgres.PostgresCFG `toml:"pg_db"`
	SRVcfg ServerCFG            `toml:"server"`
}

func (cfg *Config) InitConfig() error {

	_, err := toml.DecodeFile("./config.toml", &cfg)
	if err != nil {
		return err
	}
	return nil
}
