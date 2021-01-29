package main

import (
	"net/http"
	"time"

	"github.com/andy-smoker/wh-server/handler"
	"github.com/andy-smoker/wh-server/repository"
	"github.com/andy-smoker/wh-server/service"

	server "github.com/andy-smoker/wh-server"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	cfg := server.Config{}
	err := cfg.InitConfig()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	db, err := repository.NewPostgresDB(cfg.PGcfg)
	if err != nil {
		logrus.Panic(err)
		return
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handler.NewHandler(services)

	srv := new(Server)
	logrus.Fatal(srv.Run(cfg.SRVcfg, handler.InitRoutes()))
}

func (s *Server) Run(cfg server.ServerCFG, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              cfg.Addr + cfg.Port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20, // 1 MB
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {}
