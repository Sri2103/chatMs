package httpServer

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	userHandler "github.com/sri2103/chat_me/services/apiGateway/services/user"
	"github.com/sri2103/chat_me/services/apiGateway/utils"
)

func SetUpServer(ctx context.Context, cfg *config.Config) {
	e := echo.New()
	e.Validator = utils.SetCustomValidator()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	userHandler.RegisterUserHandler(e, cfg)
	ctxInt, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	go func() {
		if err := e.Start(cfg.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	<-ctxInt.Done()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
