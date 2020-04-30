package log

import "go.uber.org/zap"

// アプリケーションログ用のロガー
var appLogger *zap.Logger

const AppLoggerKey = "AppLogger"

func init() {
	// TODO デフォルトだと標準エラー出力に吐かれるのでログフォーマットと合わせて調整
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ := config.Build()
	appLogger = logger.Named(AppLoggerKey)
}

// getter
func GetAppLogger() *zap.Logger {
	return appLogger
}
