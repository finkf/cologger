package cologger

import (
	"testing"

	"github.com/fatih/color"
	"github.com/finkf/logger"
)

var c *color.Color

func Test(t *testing.T) {
	l := New()
	logger.Set(l)
	logger.EnableDebug(true)
	logger.Info("info message: %s", "test")
	logger.Debug("debug message: %s", "debug")
}
