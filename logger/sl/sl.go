// Package sl provides utility functions for working with the golang.org/x/exp/slog package.
//
// Example:
// Below is an example demonstrating the usage of the sl package:
//
//	package main
//
//	import (
//	    "fmt"
//	    "golang.org/x/exp/slog"
//	    "github.com/romankravchuk/logger/sl"
//	)
//
//	func main() {
//	    logger := slog.New(slog.Stderr(), slog.Debug)
//
//	    err := fmt.Errorf("some error occurred")
//
//	    logger.Log(
//	        sl.Err(err),
//	        slog.String("event", "operation_failed"),
//	        slog.Int("status_code", 500),
//	    )
//	}
package sl

import "golang.org/x/exp/slog"

// Err creates a slog attribute for logging errors with the golang.org/x/exp/slog package.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
