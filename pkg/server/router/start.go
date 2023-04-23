package router

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/handlers"
)

var (
	contestHandler *handlers.ContestHandlers
)

func initServer() {
	contestRepository := inmemory.NewContestRepository([]domain.Contest{})
	contestHandler = handlers.NewContestHandlers(
		*controller.NewContestController(
			contestRepository,
			*contest.NewCreateContestService(contestRepository),
		),
	)
}

func StartServer(port int) {
	initServer()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.HideBanner = true

	// routerの呼び出し
	rootRouter(e)

	// グレイスフルシャットダウン用
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
			e.Logger.Fatal("Shutting down server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}