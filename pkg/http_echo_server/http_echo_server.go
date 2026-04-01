package http_echo_server

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	Echo   *echo.Echo
	listen string
}

func New(listen string, middlewares ...func(http.Handler) http.Handler) Server {
	e := echo.New()
	e.Use(middleware.RequestLogger()) // use the RequestLogger middleware with slog logger
	e.Use(middleware.Recover())       // recover panics as errors for proper error handling

	return Server{
		Echo:   e,
		listen: listen,
	}
}

func (s *Server) Start() chan error {
	errCh := make(chan error)
	go func() {
		errCh <- s.Echo.Start(s.listen)
		close(errCh)
	}()
	return errCh
}

func (s *Server) Stop() {

}
