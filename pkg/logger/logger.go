package logger

// DefaultLogger は デフォルトの logger を提供します。
var DefaultLogger = New()

// LevelLogger は ログレベルの設定された logger の interface です。
type LevelLogger interface {
	Tracef(ormat string, args ...interface{})
	Debugf(ormat string, args ...interface{})
	Infof(ormat string, args ...interface{})
	Warnf(ormat string, args ...interface{})
	Errorf(ormat string, args ...interface{})
}

// interfaces
var _ LevelLogger = (*Logger)(nil)

// Tracef は TRACE レベルのログを出力します。
func Tracef(format string, args ...interface{}) { DefaultLogger.Tracef(format, args...) }

// Debugf は DEBUG レベルのログを出力します。
func Debugf(format string, args ...interface{}) { DefaultLogger.Debugf(format, args...) }

// Infof は INFO レベルのログを出力します。
func Infof(format string, args ...interface{}) { DefaultLogger.Infof(format, args...) }

// Warnf は WARN レベルのログを出力します。
func Warnf(format string, args ...interface{}) { DefaultLogger.Warnf(format, args...) }

// Errorf は ERROR レベルのログを出力します。
func Errorf(format string, args ...interface{}) { DefaultLogger.Errorf(format, args...) }
