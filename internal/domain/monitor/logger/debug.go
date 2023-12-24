package logger

type Debug interface {
	DebugF(message string, args ...any)
	DebugKV(message string, args ...any)
}
