package logger

import "context"

type WithContext interface {
	WithContext(ctx context.Context) Logger
}
