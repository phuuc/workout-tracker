package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/finnpn/workout-tracker/pkg/helpers"
)

type level int8

const (
	levelInfo level = iota
	levelError
)

func (l level) String() string {
	switch l {
	case levelError:
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
	Level   string         `json:"level"`
	Message interface{}    `json:"message"`
	Trace   *helpers.Trace `json:"trace,omitempty"`
	Time    string         `json:"time"`
}

func NewLog() *Log {
	return &Log{
		out: os.Stdout,
	}
}

func (l *Log) Info(message string, a ...any) (int, error) {
	return l.print(levelInfo, message, a...)
}
func (l *Log) Error(message string, a ...any) (int, error) {
	return l.print(levelError, message, a...)
}

func (l *Log) print(level level, message string, a ...any) (int, error) {
	message = fmt.Sprintf(message, a...)
	logBody := &LogBody{
		Level:   level.String(),
		Message: message,
		Time:    time.Now().UTC().Format(time.RFC3339),
	}
	if level >= levelError {
		logBody.Trace = helpers.NewTrace()
	}

	line, err := json.Marshal(logBody)
	if err != nil {
		return 0, errors.New("unable to marshal log message")
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if level >= levelError {
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
