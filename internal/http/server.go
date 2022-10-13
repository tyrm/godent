package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
)

const serverTimeout = 60 * time.Second

// Server is a http 2 web server.
type Server struct {
	router *mux.Router
	srv    *http.Server
}

// NewServer creates a new http web server.
func NewServer(_ context.Context) (*Server, error) {
	r := mux.NewRouter()

	s := &http.Server{
		Addr:         viper.GetString(config.Keys.ServerHTTPBind),
		Handler:      r,
		WriteTimeout: serverTimeout,
		ReadTimeout:  serverTimeout,
	}

	server := &Server{
		router: r,
		srv:    s,
	}

	// add global middlewares
	r.Use(server.WrapInMiddlewares)

	r.NotFoundHandler = server.NotFoundHandler()
	r.MethodNotAllowedHandler = server.MethodNotAllowedHandler()

	return server, nil
}

// HandleFunc attaches a function to a path.
func (s *Server) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f)
}

// PathPrefix attaches a new route url path prefix.
func (s *Server) PathPrefix(path string) *mux.Route {
	return s.router.PathPrefix(path)
}

// Start starts the web server.
func (s *Server) Start() error {
	l := logger.WithField("func", "Start")
	l.Infof("listening on %s", s.srv.Addr)

	return s.srv.ListenAndServe()
}

// Stop shuts down the web server.
func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) MethodNotAllowedHandler() http.Handler {
	// wrap in middleware since middleware isn't run on error pages
	return s.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReturnError(w, ErrCodeMethodNotAllowed, r.Method)
	}))
}

func (s *Server) NotFoundHandler() http.Handler {
	// wrap in middleware since middleware isn't run on error pages
	return s.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ReturnError(w, ErrCodeNotFound, r.URL.Path)
	}))
}
