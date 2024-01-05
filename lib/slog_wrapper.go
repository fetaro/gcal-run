package lib

import (
	"fmt"
	"log/slog"
	"os"
)

type SLogWrapper struct {
	logger *slog.Logger
}

func NewSLogWrapper(logger *slog.Logger) *SLogWrapper {
	return &SLogWrapper{
		logger: logger,
	}
}

func (s *SLogWrapper) Debug(format string, a ...any) {
	s.logger.Debug(fmt.Sprintf(format, a...))
}
func (s *SLogWrapper) Info(format string, a ...any) {
	s.logger.Info(fmt.Sprintf(format, a...))
}
func (s *SLogWrapper) Warn(format string, a ...any) {
	s.logger.Warn(fmt.Sprintf(format, a...))
}
func (s *SLogWrapper) Error(format string, a ...any) {
	s.logger.Error(fmt.Sprintf(format, a...))
}

var loggerSingleton *SLogWrapper

func GetLogger() *SLogWrapper {
	isDebug := os.Getenv("DEBUG")
	var logLevel slog.Level
	if isDebug == "" {
		logLevel = slog.LevelInfo
	} else {
		logLevel = slog.LevelDebug
	}
	if loggerSingleton != nil {
		return loggerSingleton
	} else {
		handler := NewSimpleHandler(os.Stdout, logLevel)
		l := NewSLogWrapper(slog.New(handler))
		loggerSingleton = l
		return l
	}
}
