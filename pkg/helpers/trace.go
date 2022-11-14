package helpers

import (
	"runtime"
)

type Trace struct {
	FileName string `json:"file_name"`
	FuncName string `json:"func_name"`
	Line     int    `json:"line"`
}

func NewTrace() *Trace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return &Trace{
		FileName: frame.File,
		FuncName: frame.Function,
		Line:     frame.Line,
	}
}
