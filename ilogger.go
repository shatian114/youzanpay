package youzan

import "context"

// 设置外部日志记录器
type ILogger interface {
	Error(ctx context.Context, v ...any)
	Info(ctx context.Context, v ...any)
	Fatal(ctx context.Context, v ...any)
	Debug(ctx context.Context, v ...any)
}

var Logger ILogger

func Debug(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Debug(ctx, v)
	}
}

func Info(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Info(ctx, v)
	}
}

func Error(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Error(ctx, v)
	}
}

func Fatal(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Fatal(ctx, v)
	}
}

func SetLogger(l ILogger) {
	Logger = l
}
