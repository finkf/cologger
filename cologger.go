package cologger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/finkf/logger"
)

// Default colors for info and debug messages and the time format string.
var (
	DefaultInfoColor                = color.New(color.FgCyan)
	DefaultDebugColor               = color.New(color.FgMagenta)
	DefaultTimeColor                = color.New(color.FgYellow)
	_                 logger.Logger = &Logger{}
)

// WithDebugColor sets the color of the debug tag.
func WithDebugColor(c *color.Color) func(*Logger) {
	return func(l *Logger) {
		l.dColor = c
	}
}

// WithTimeColor sets the color of the time string.
func WithTimeColor(c *color.Color) func(*Logger) {
	return func(l *Logger) {
		l.tColor = c
	}
}

// WithInfoColor sets the color of the info tag.
func WithInfoColor(c *color.Color) func(*Logger) {
	return func(l *Logger) {
		l.iColor = c
	}
}

// WithTimeFormat sets the time format string.
func WithTimeFormat(tfmt string) func(*Logger) {
	return func(l *Logger) {
		l.tfmt = tfmt
	}
}

// WithWriter sets the output writer.
func WithWriter(out io.Writer) func(*Logger) {
	return func(l *Logger) {
		l.out = out
	}
}

// Logger implements logger.Logger with colors.
type Logger struct {
	out                    io.Writer
	tfmt                   string
	dColor, iColor, tColor *color.Color
}

// New creates a Logger instance using stderr and the default settings.
func New(args ...func(*Logger)) *Logger {
	l := &Logger{
		out:    os.Stderr,
		tfmt:   logger.DefaultTimeFormat,
		dColor: DefaultDebugColor,
		iColor: DefaultInfoColor,
		tColor: DefaultTimeColor,
	}
	for _, arg := range args {
		arg(l)
	}
	return l
}

// Debug writes a debug message if it is enabled.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.write(l.dColor.FprintfFunc(), "DEBUG", format, args...)
}

// Info writes a debug message.
func (l *Logger) Info(format string, args ...interface{}) {
	l.write(l.iColor.FprintfFunc(), "INFO", format, args...)
}

type fpfunc func(io.Writer, string, ...interface{})

func (l *Logger) write(p fpfunc, kind, format string, args ...interface{}) {
	l.tColor.Fprintf(l.out, time.Now().Format(l.tfmt))
	fmt.Fprint(l.out, " ")
	p(l.out, kind)
	fmt.Fprint(l.out, " ")
	fmt.Fprintf(l.out, format, args...)
	fmt.Fprintln(l.out)
}
