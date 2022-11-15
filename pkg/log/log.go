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

var Log *Logger

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

type Logger struct {
	out io.Writer
	mu  sync.Mutex
}

type LogBody struct {
	Level   string         `json:"level"`
	Message interface{}    `json:"message"`
	Trace   *helpers.Trace `json:"trace,omitempty"`
	Time    string         `json:"time"`
}

func init() {
	Log = NewLog()
}

func NewLog() *Logger {
	return &Logger{
		out: os.Stdout,
	}
}

func Info(message string, a ...any) {
	print(levelInfo, message, a...)
}

func Error(message string, a ...any) {
	print(levelError, message, a...)
}

func print(level level, message string, a ...any) (int, error) {
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
	Log.mu.Lock()
	defer Log.mu.Unlock()
	if level >= levelError {
		return printError(line)
	}
	return printInfo(line)
}

func printInfo(line []byte) (int, error) {
	return Log.out.Write(append([]byte(line), '\n'))
}

func printError(line []byte) (int, error) {
	var (
		reset = []byte("\033[0m")
		red   = []byte("\033[31m")
		body  = make([]byte, 0)
	)
	body = append(body, red...)
	body = append(body, line...)
	body = append(body, reset...)

	return Log.out.Write(append(body, '\n'))
}
