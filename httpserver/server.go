package httpserver

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultAddr            = ":80"
	defaultShutdownTimeout = 3 * time.Second
)

var defaultErrorLogger = slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo)

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

// New creates a new http server.
func New(ctx context.Context, handler http.Handler, opts ...Option) *Server {
	srv := &Server{
		server: &http.Server{
			Handler:      handler,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Addr:         defaultAddr,
			BaseContext: func(_ net.Listener) context.Context {
				return ctx
			},
			ErrorLog: defaultErrorLogger,
		},
		shutdownTimeout: defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

// Run starts the server. If context is cancelled, the server will be gracefully shutdown.
func (s *Server) Run(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return s.server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		return s.server.Shutdown(ctx)
	})

	if err := g.Wait(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
