package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	With(keyvals ...any) Logger
	Info(msg string, keyvals ...any)
	Error(msg string, keyvals ...any)
}

type stdLogger struct {
	prefix string
	l      *log.Logger
}

func NewStdLogger() Logger {
	return &stdLogger{
		l: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (s *stdLogger) With(keyvals ...any) Logger {
	p := s.prefix
	for i := 0; i+1 < len(keyvals); i += 2 {
		p += " " + kvString(keyvals[i], keyvals[i+1])
	}
	return &stdLogger{prefix: p, l: s.l}
}

func (s *stdLogger) Info(msg string, keyvals ...any) {
	line := "[INFO] " + s.prefix + " " + msg
	for i := 0; i+1 < len(keyvals); i += 2 {
		line += " " + kvString(keyvals[i], keyvals[i+1])
	}
	s.l.Output(2, line)
}

func (s *stdLogger) Error(msg string, keyvals ...any) {
	line := "[ERROR] " + s.prefix + " " + msg
	for i := 0; i+1 < len(keyvals); i += 2 {
		line += " " + kvString(keyvals[i], keyvals[i+1])
	}
	s.l.Output(2, line)
}

func kvString(k, v any) string {
	return "[" + toString(k) + "=" + toString(v) + "]"
}

func toString(v any) string {
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmtSprint(t)
	}
}

// small fmt helper to avoid importing fmt in hot path repeatedly
func fmtSprint(v any) string { return fmt.Sprintf("%v", v) }
