package log

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
	"time"

	"github.com/finnpn/workout-tracker/pkg/trace"
)

type level int8

const (
	LevelInfo level = iota
	LevelError
)

func (l level) String() string {
	switch l {
	case LevelError:
		return "ERROR"
	default:
		return "INFO"
	}
}

type Log struct {
	out io.Writer
	mu  sync.Mutex
}

type LogBody struct {
	Level   string       `json:"level"`
	Message string       `json:"message"`
	Trace   *trace.Trace `json:"trace,omitempty"`
	Time    string       `json:"time"`
}

func NewLog() *Log {
	return &Log{
		out: os.Stdout,
	}
}

func (l *Log) Info(message string) (int, error) {
	return l.print(LevelInfo, message)
}
func (l *Log) Error(message string) (int, error) {
	return l.print(LevelError, message)
}

func (l *Log) print(level level, message string) (int, error) {
	logBody := &LogBody{
		Level:   level.String(),
		Message: message,
		Time:    time.Now().UTC().Format(time.RFC3339),
	}
	if level >= LevelError {
		logBody.Trace = trace.NewTrace()
	}
	line, err := json.Marshal(logBody)
	if err != nil {
		return 0, errors.New("unable to marshal log message")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if level >= LevelError {
		return l.printError(line)
	}
	return l.printInfo(line)
}

func (l *Log) printInfo(line []byte) (int, error) {
	return l.out.Write(append([]byte(line), '\n'))
}

func (l *Log) printError(line []byte) (int, error) {
	var (
		reset = []byte("\033[0m")
		red   = []byte("\033[31m")
		body  = make([]byte, 0)
	)
	body = append(body, red...)
	body = append(body, line...)
	body = append(body, reset...)

	return l.out.Write(append(body, '\n'))
}
