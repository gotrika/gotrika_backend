package bootstrap

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	transport "github.com/gotrika/gotrika_backend/internal/transport/http"
	"github.com/gotrika/gotrika_backend/internal/transport/http/handler"
	"github.com/gotrika/gotrika_backend/pkg/logger"
)

// RunHTTP initializes http aplication.
// @securitydefinitions.oauth2.password ApiAuth
// @tokenUrl /api/auth/sign-in/
// @description api auth
// @in header
// @name Authorization
func RunHTTP() {
	deps, err := InitDependencies()
	if err != nil {
		logger.Error("Invalid get dependenices")
		return
	}
	handlers := handler.NewHandler(deps.Services())
	server := transport.NewHTTPServer(deps.Config(), handlers.Init(deps.Config()))
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := deps.MongoClient().Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
	}
}
