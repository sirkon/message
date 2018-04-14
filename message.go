package message

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"path/filepath"
)

// Color ...
type Color string

const (
	reset   Color = "\033[0m"
	black         = "\033[30m"
	boldred       = "\033[1;31m"
	red           = "\033[31m"
	green         = "\033[32m"
	yellow        = "\033[33m"
	blue          = "\033[34m"
	magent        = "\033[35m"
	cyan          = "\033[36m"
	white         = "\033[37m"
)

const (
	anchor = "\033["
)

var debug bool
var useColor bool = true

func split(message string) []string {
	var splits []string
	ci := ColorInfo{}
	for len(message) > 0 {
		ok, _ := ci.Extract(message)
		if !ok {
			splits = append(splits, message)
			break
		}
		if len(ci.Text) > 0 {
			splits = append(splits, ci.Text)
		}
		if useColor {
			splits = append(splits, fmt.Sprintf("\033[%dm", ci.Code))
		}
		message = ci.rest
	}
	return splits
}

// fileLine retuns string "<file name>:<line number>". File name and line number are the first that are not from this
// package
func fileLine() string {
	fpcs := make([]uintptr, 10)
	frameIter := runtime.CallersFrames(fpcs[:runtime.Callers(2, fpcs)])
	for {
		frame, more := frameIter.Next()
		_, name := filepath.Split(frame.Function)
		if !strings.HasPrefix(name, "message.") {
			return fmt.Sprintf("%s:%d", frame.File, frame.Line)
		}
		if !more {
			return "<unknown file>:<unknown line>"
		}
	}
}

func printer(color Color, message string) string {
	if debug {
		message = fileLine() + " " + message
	}
	parts := split(message)
	buffer := &bytes.Buffer{}
	if useColor {
		buffer.WriteString(string(color))
	}
	for _, p := range parts {
		buffer.WriteString(p)
		if p == string(reset) && useColor {
			buffer.WriteString(string(color))
		}
	}
	if useColor {
		buffer.WriteString(string(reset))
	}
	return buffer.String()
}

func printLine(color Color, data ...interface{}) {
	os.Stderr.Write([]byte(printer(color, fmt.Sprintln(data...))))
}

func printFormat(color Color, format string, data ...interface{}) {
	msg := fmt.Sprintf(format, data...)
	printLine(color, msg)
}

// Fatalf print yellow formatted message
func Fatalf(format string, data ...interface{}) {
	printFormat(boldred, format, data...)
	os.Exit(1)
}

// Fatal print red message and exit
func Fatal(data ...interface{}) {
	printLine(boldred, data...)
	os.Exit(1)
}

// Criticalf print yellow formatted message
func Criticalf(format string, data ...interface{}) {
	printFormat(boldred, format, data...)
	os.Exit(1)
}

// Critical print red message and exit
func Critical(data ...interface{}) {
	printLine(boldred, data...)
	os.Exit(1)
}

// Errorf print yellow formatted message
func Errorf(format string, data ...interface{}) {
	printFormat(boldred, format, data...)
}

// Error print red message and exit
func Error(data ...interface{}) {
	printLine(boldred, data...)
}

// Warningf print yellow formatted message
func Warningf(format string, data ...interface{}) {
	printFormat(yellow, format, data...)
}

// Warning print yellow message
func Warning(data ...interface{}) {
	printLine(yellow, data...)
}

// Noticef print yellow formatted message
func Noticef(format string, data ...interface{}) {
	printFormat(green, format, data...)
}

// Notice print yellow message
func Notice(data ...interface{}) {
	printLine(green, data...)
}

// Infof print yellow formatted message
func Infof(format string, data ...interface{}) {
	printFormat(reset, format, data...)
}

// Info print yellow message
func Info(data ...interface{}) {
	printLine(reset, data...)
}

// Debugf print yellow formatted message
func Debugf(format string, data ...interface{}) {
	printFormat(cyan, format, data...)
}

// Debug print yellow message
func Debug(data ...interface{}) {
	printLine(cyan, data...)
}

// SetDebug sets debug on/off depending on the status
func SetDebug(status bool) {
	debug = status
}

// UseColor enables/disables color output depending on the decision's value
func UseColor(decision bool) {
	useColor = decision
}

func init() {
	if os.Getenv("DEBUG") != "" {
		debug = true
	}
}
