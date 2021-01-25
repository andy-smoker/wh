package main

import (
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	server "github.com/andy-smoker/wh-server"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := server.Config{}
	err := initConfig()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	srv := &http.Server{
		Addr: cfg.SRVcfg.Addr + cfg.SRVcfg.Port,
		//Handler:           handlers.InitRoutes(),
		MaxHeaderBytes:    1 << 20, // 1 MB
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}

func initConfig() error {
	cfg := server.Config{}

	_, err := toml.DecodeFile("./config.toml", &cfg)
	if err != nil {
		return err
	}
	return nil
}
