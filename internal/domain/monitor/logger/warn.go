package logger

type Warn interface {
	WarnF(message string, args ...any)
	WarnKV(message string, args ...any)
}
