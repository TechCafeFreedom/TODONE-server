package log

import "go.uber.org/zap"

// アプリケーションログ用のロガー
var appLogger *zap.Logger

const AppLoggerKey = "AppLogger"

func init() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	config.DisableStacktrace = true // スタックトレースONにしたい場合はfalseにする
	logger, _ := config.Build()
	appLogger = logger.Named(AppLoggerKey)
}

// getter
func GetAppLogger() *zap.Logger {
	return appLogger
}
