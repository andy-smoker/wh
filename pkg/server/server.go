package server

import (
	"net/http"
	"time"

	"github.com/andy-smoker/wh-server/pkg/config"
	"github.com/andy-smoker/wh-server/pkg/handler"
	"github.com/andy-smoker/wh-server/pkg/repository"
	"github.com/andy-smoker/wh-server/pkg/repository/postgres"
	"github.com/andy-smoker/wh-server/pkg/service"
	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func Run() {
	cfg := config.Config{}
	err := cfg.InitConfig()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	db, err := postgres.NewDB(cfg.PGcfg)
	if err != nil {
		logrus.Panic(err)
		return
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handler.NewHandler(services)

	srv := new(Server)
	logrus.Fatal(srv.NewServer(cfg.SRVcfg, handler.InitRoutes()))
}

func (s *Server) NewServer(cfg config.ServerCFG, handler http.Handler) error {
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
