// start name:top
package main

// start name:imports type:merge
import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/zlietapki/boilerplate/internal/config"
	"github.com/zlietapki/boilerplate/internal/repository"
	"github.com/zlietapki/boilerplate/internal/rest_handler"
	"github.com/zlietapki/boilerplate/internal/usecase"
	"github.com/zlietapki/boilerplate/pkg/httpserver"
)

// start name:main
func main() {
	cfg := config.New()

	repo := repository.New()
	uc := usecase.New(repo)

	//start name:handler type:add
	restHandler := rest_handler.New(uc)

	httpServer := httpserver.New(cfg.HTTPListen)

	api := httpServer.Srv.Group("/api")
	api.GET("/users", restHandler.GetUsers)
	api.POST("/users", restHandler.CreateUser)

	httpErrCh := httpServer.Start()
	slog.Info("HTTP listening", "addr", cfg.HTTPListen)

	//start name:signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:

	//start name:signals_handler type:add
	case err := <-httpErrCh:
		panic("http server" + err.Error())

		// start name:bottom
	}
}
