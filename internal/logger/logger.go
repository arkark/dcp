package logger

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	NONE
)

func (level LogLevel) toString() string {
	return []string{
		"DEBUG", "INFO", "WARN", "ERROR",
	}[level]
}
