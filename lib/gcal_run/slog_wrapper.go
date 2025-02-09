package gcal_run

import (
	"fmt"
	"io"
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

func GetLogger(logFilePath string) (*SLogWrapper, error) {
	isDebug := os.Getenv("DEBUG")
	var logLevel slog.Level
	if isDebug == "" {
		logLevel = slog.LevelInfo
	} else {
		logLevel = slog.LevelDebug
	}
	if loggerSingleton != nil {
		return loggerSingleton, nil
	} else {
		var writer io.Writer
		if logFilePath != "" {
			logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				return nil, err
			}
			writer = io.MultiWriter(os.Stdout, logFile)
		} else {
			writer = os.Stdout
		}
		handler := NewSimpleHandler(writer, logLevel)
		l := NewSLogWrapper(slog.New(handler))
		loggerSingleton = l
		return l, nil
	}
}
