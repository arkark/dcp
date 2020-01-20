// +build debug

package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var minLevel LogLevel

func Init(path string, level LogLevel) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags)
	minLevel = level
	return nil
}

func Write(level LogLevel, format string, v ...interface{}) {
	if level < minLevel {
		return
	}
	log.SetPrefix(
		fmt.Sprintf(
			"[%s] %s ",
			level.toString(),
			where(),
		),
	)
	log.Println(fmt.Sprintf(format, v...))
}

func where() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}

	return fmt.Sprintf(
		"%s:%d",
		file,
		line,
	)
}
