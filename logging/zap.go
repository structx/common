// Package logging initialize third party loggers
package logging

import (
	"fmt"

	"go.uber.org/zap"
)

// NewZap creates a new zap logger
func NewZap() (*zap.Logger, error) {

	l, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create zap logger: %w", err)
	}

	return l, nil
}
