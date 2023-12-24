package monitor

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/MarlakDevelop/hotel-booking/internal/domain/monitor/logger"
)

const (
	debugLevelName = "DEBUG"
	infoLevelName  = "INFO"
	warnLevelName  = "WARN"
	errorLevelName = "ERROR"

	requestIDKey = "requestID"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(writer io.Writer, levelName string) *Logger {
	var level slog.Level

	switch levelName {
	case debugLevelName:
		level = slog.LevelDebug
	case infoLevelName:
		level = slog.LevelInfo
	case warnLevelName:
		level = slog.LevelWarn
	case errorLevelName:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	return &Logger{logger: slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level}))}
}

func (log *Logger) WithContext(ctx context.Context) logger.Logger {
	args := make([]any, 0)

	requestID := middleware.GetReqID(ctx)
	if requestID != "" {
		args = append(args, requestIDKey, requestID)
	}

	return &Logger{logger: log.logger.With(args...)}
}

func (log *Logger) DebugF(message string, args ...any) {
	log.logger.Debug(fmt.Sprintf(message, args...))
}

func (log *Logger) DebugKV(message string, args ...any) {
	log.logger.Debug(message, args...)
}

func (log *Logger) InfoF(message string, args ...any) {
	log.logger.Info(fmt.Sprintf(message, args...))
}

func (log *Logger) InfoKV(message string, args ...any) {
	log.logger.Info(message, args...)
}

func (log *Logger) WarnF(message string, args ...any) {
	log.logger.Warn(fmt.Sprintf(message, args...))
}

func (log *Logger) WarnKV(message string, args ...any) {
	log.logger.Warn(message, args...)
}

func (log *Logger) ErrorF(message string, args ...any) {
	log.logger.Error(fmt.Sprintf(message, args...))
}

func (log *Logger) ErrorKV(message string, args ...any) {
	log.logger.Error(message, args...)
}
