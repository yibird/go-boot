package initialize

import (
	"context"
	"errors"
	"fmt"
	"go-boot/config"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer(ew *EngineWrapper, config *config.Config, logger *zap.Logger) {
	engine := ew.Engine
	application := config.Application
	addr := fmt.Sprintf(":%s", config.System.Addr)
	logger.Info(fmt.Sprintf("💩%s started successfully🎉🎉🎉, port%s, version: %s", application.Name, addr, application.Version))
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(fmt.Sprintf("listen: %s\n", err.Error()))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(fmt.Sprintf("Server forced to shutdown: %s", err.Error()))
	}
	logger.Info("server exiting")
}
