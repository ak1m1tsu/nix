// Package log provides logging functionality.
package log

import (
	"io"
	"log/slog"
)

// Env represents the environment in which the application is running.
type Env string

const (
	Development Env = "development"
	Local       Env = "local"
	Production  Env = "production"
)

// New creates a new logger using the specified writer and environment.
//
// For Development returns JSON logger with DEBUG level.
// For Local returns Text logger with DEBUG level.
// For Production returns JSON logger with INFO level.
func New(w io.Writer, env Env) *slog.Logger {
	var (
		s       *slog.Logger
		l, json = envToOpts(env)
		opts    = &slog.HandlerOptions{Level: l}
	)

	if json {
		s = slog.New(slog.NewJSONHandler(w, opts))
	} else {
		s = slog.New(slog.NewTextHandler(w, opts))
	}

	return s
}

// envToOpts converts the environment to logger options.
func envToOpts(env Env) (level slog.Level, json bool) {
	switch env {
	case Development:
		level = slog.LevelDebug
		json = true
	case Local:
		level = slog.LevelDebug
	case Production:
		json = true
	default:
		return
	}
	return
}
