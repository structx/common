package logging_test

import (
	"testing"

	"github.com/structx/common/logging"
)

func TestNew(t *testing.T) {

	logger, err := logging.NewZap()
	if err != nil {
		t.Error(err)
	}

	logger.Info("test")
}
