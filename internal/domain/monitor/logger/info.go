package logger

type Info interface {
	InfoF(message string, args ...any)
	InfoKV(message string, args ...any)
}
