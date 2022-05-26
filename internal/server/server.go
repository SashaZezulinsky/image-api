package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/image-api/config"
	"github.com/image-api/pkg/logger"
	"github.com/labstack/echo/v4"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	echo    *echo.Echo
	cfg     *config.Config
	mongoDB *mongo.Client
	logger  logger.Logger
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, mongoDB *mongo.Client, logger logger.Logger) *Server {
	return &Server{echo: echo.New(), cfg: cfg, mongoDB: mongoDB, logger: logger}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	s.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		s.logger.Errorw("Error on request", "Path", c.Path(), "Params", c.QueryParams(), "Err", err)
		s.echo.DefaultHTTPErrorHandler(err, c)
	}

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
