package httpServer

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sri2103/chat_me/services/apiGateway/config"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func SetCustomValidator(e *echo.Echo) *echo.Echo {
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}
	return e
}

func SetUpServer(ctx context.Context, cfg *config.Config) {
	e := echo.New()
	e = SetCustomValidator(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
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
