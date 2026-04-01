// start name:top
package main

// start name:imports type:merge
import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/zlietapki/gena/internal/config"
	"github.com/zlietapki/gena/internal/map_repo"
	"github.com/zlietapki/gena/internal/rest_handler"
	"github.com/zlietapki/gena/internal/usecase"
	"github.com/zlietapki/gena/pkg/http_echo_server"
)

// start name:main
func main() {
	cfg := config.New()

	// start name:usecase_deps type:add
	mapRepo := map_repo.New()

	// start name:new_usecase
	uc := usecase.New(usecase.Depends{
		// start name:usecase_deps_objs type:add
		MemRepo: mapRepo,
		// start name:post_usecase_deps_objs
	})

	//start name:after_usecase type:add
	restHandler := rest_handler.New(uc)

	httpServer := http_echo_server.New(cfg.HTTPListen)

	api := httpServer.Echo.Group("/api")
	api.GET("/counter", restHandler.GetCounter)

	httpErrCh := httpServer.Start()
	slog.Info("HTTP listening", "addr", cfg.HTTPListen)

	//start name:signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:

	//start name:signals_select type:add
	case err := <-httpErrCh:
		panic("http server" + err.Error())
		//start name:post_signals_select
	}
	// start name:bottom
}
