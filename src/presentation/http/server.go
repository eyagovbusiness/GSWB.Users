package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eyagovbusiness/GSWB.Users/pkg/logger"
	"github.com/eyagovbusiness/GSWB.Users/src/presentation/http/handler"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(userHandler *handler.UserHandler) *Server {
	logger.Init()

	engine := SetupRouter(userHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	return &Server{httpServer: srv}
}

func (s *Server) Start(ctx context.Context) error {
	logger.Logger.Info("Starting HTTP server on :8080")

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("HTTP server error", "error", err)
		}
	}()

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	logger.Logger.Info("Shutting down HTTP server...")

	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctxTimeout)
}
