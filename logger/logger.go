package logger

import "log"

type Logger struct {
	enabled bool
}

func InitLogger(enabled bool) Logger {
	return Logger{
		enabled,
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.enabled {
		log.Printf(format, v...)
	}
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
