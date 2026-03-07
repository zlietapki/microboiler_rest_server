package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

//type addHandlerFunc func(s *grpc.Server)

type Server struct {
	Srv    *echo.Echo
	listen string
}

func New(listen string, middlewares ...func(http.Handler) http.Handler) Server {
	e := echo.New()
	e.Use(middleware.RequestLogger()) // use the RequestLogger middleware with slog logger
	e.Use(middleware.Recover())       // recover panics as errors for proper error handling

	return Server{
		Srv:    e,
		listen: listen,
	}
}

func (s *Server) Start() chan error {
	errCh := make(chan error)
	go func() {
		errCh <- s.Srv.Start(s.listen)
		close(errCh)
	}()
	return errCh
}

func (s *Server) Stop() {

}
