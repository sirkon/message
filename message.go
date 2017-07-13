package message

import (
	"bytes"
	"fmt"
	"os"
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

func split(message string) []string {
	var splits []string
	for len(message) > 0 {
		p := &colorExtractor{}
		ok, _ := p.Parse([]byte(message))
		if !ok {
			splits = append(splits, message)
			break
		}
		if len(p.Head) > 0 {
			splits = append(splits, string(p.Head))
		}
		splits = append(splits, "\033["+string(p.Code)+"m")
		message = string(p.Rest)
	}
	return splits
}

func printer(color Color, message string) string {
	parts := split(message)
	buffer := &bytes.Buffer{}
	buffer.WriteString(string(color))
	for _, p := range parts {
		buffer.WriteString(p)
		if p == string(reset) {
			buffer.WriteString(string(color))
		}
	}
	buffer.WriteString(string(reset))
	return buffer.String()
}

func println(color Color, data ...interface{}) {
	os.Stderr.Write([]byte(printer(color, fmt.Sprintln(data...))))
}

func printf(color Color, format string, data ...interface{}) {
	msg := fmt.Sprintf(format, data...)
	println(color, msg)
}

// Criticalf print yellow formatted message
func Criticalf(format string, data ...interface{}) {
	printf(boldred, format, data...)
	os.Exit(1)
}

// Critical print red message and exit
func Critical(data ...interface{}) {
	println(boldred, data...)
	os.Exit(1)
}

// Errorf print yellow formatted message
func Errorf(format string, data ...interface{}) {
	printf(boldred, format, data...)
}

// Error print red message and exit
func Error(data ...interface{}) {
	println(boldred, data...)
}

// Warningf print yellow formatted message
func Warningf(format string, data ...interface{}) {
	printf(yellow, format, data...)
}

// Warning print yellow message
func Warning(data ...interface{}) {
	println(yellow, data...)
}

// Noticef print yellow formatted message
func Noticef(format string, data ...interface{}) {
	printf(green, format, data...)
}

// Notice print yellow message
func Notice(data ...interface{}) {
	println(green, data...)
}

// Infof print yellow formatted message
func Infof(format string, data ...interface{}) {
	printf(reset, format, data...)
}

// Info print yellow message
func Info(data ...interface{}) {
	println(reset, data...)
}
