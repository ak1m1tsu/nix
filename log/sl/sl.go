// Package sl provides utility functions for working with the log/slog package.
//
// Example:
// Below is an example demonstrating the usage of the sl package:
//
//	package main
//
//	import (
//	    "fmt"
//	    "log/slog"
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

import "log/slog"

// Err creates a slog attribute for logging errors with the log/slog package.
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
