// +build !debug

package logger

func Init(path string, level LogLevel) error {
	return nil
}

func Write(level LogLevel, format string, v ...interface{}) {
}
