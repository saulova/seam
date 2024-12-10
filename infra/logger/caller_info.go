package logger

import (
	"runtime"
)

type CallerInfo struct {
	Caller string
	File   string
	Line   int
}

func NewCallerInfo(skip int) *CallerInfo {
	pc, file, line, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return &CallerInfo{
			Caller: details.Name(),
			File:   file,
			Line:   line,
		}
	}

	return &CallerInfo{
		Caller: "unknown",
		File:   "unknown",
		Line:   0,
	}
}
