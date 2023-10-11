// Package logging initialize third party loggers
package logging

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// NewZap creates a new zap logger
func NewZap() (*zap.Logger, error) {

	debug := strings.ToLower(os.Getenv("DEBUG_LEVEL"))
	if debug == "production" {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}
