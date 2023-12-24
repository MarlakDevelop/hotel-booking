package logger

type Error interface {
	ErrorF(message string, args ...any)
	ErrorKV(message string, args ...any)
}
