package logger

import (
	"os"
	"time"

	zerolog "github.com/rs/zerolog"
)

// Logger は logger の interface です。
type Logger struct {
	Zerolog zerolog.Logger
}

// New は logger を生成します。
func New() Logger {
	w := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	}

	logger := Logger{
		Zerolog: zerolog.New(w).With().Timestamp().Logger(),
	}

	return logger
}

// Tracef は TRACE レベルのログを出力します。
func (l *Logger) Tracef(format string, args ...interface{}) {
	l.Zerolog.Trace().Msgf(format, args...)
}

// Debugf は DEBUG レベルのログを出力します。
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Zerolog.Debug().Msgf(format, args...)
}

// Infof は INFO レベルのログを出力します。
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Zerolog.Info().Msgf(format, args...)
}

// Warnf は WARN レベルのログを出力します。
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Zerolog.Warn().Msgf(format, args...)
}

// Errorf は ERROR レベルのログを出力します。
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Zerolog.Error().Msgf(format, args...)
}
