package message

import (
	"bytes"
	"fmt"
	"io"
)

type cons struct {
	dst io.Writer
}

func (c cons) print(name string, value any) {
	var buf bytes.Buffer
	buf.WriteString("    ")
	if useColor {
		buf.WriteString("\033[1m")
	}
	buf.WriteString(name)
	if useColor {
		buf.WriteString("\033[0m")
	}
	buf.WriteString(": ")
	fmt.Fprintln(&buf, value)
	io.Copy(c.dst, &buf)
}

func (c cons) Bool(name string, value bool) {
	c.print(name, value)
}

func (c cons) Int(name string, value int) {
	c.print(name, value)
}

func (c cons) Int8(name string, value int8) {
	c.print(name, value)
}

func (c cons) Int16(name string, value int16) {
	c.print(name, value)
}

func (c cons) Int32(name string, value int32) {
	c.print(name, value)
}

func (c cons) Int64(name string, value int64) {
	c.print(name, value)
}

func (c cons) Uint(name string, value uint) {
	c.print(name, value)
}

func (c cons) Uint8(name string, value uint8) {
	c.print(name, value)
}

func (c cons) Uint16(name string, value uint16) {
	c.print(name, value)
}

func (c cons) Uint32(name string, value uint32) {
	c.print(name, value)
}

func (c cons) Uint64(name string, value uint64) {
	c.print(name, value)
}

func (c cons) Float32(name string, value float32) {
	c.print(name, value)
}

func (c cons) Float64(name string, value float64) {
	c.print(name, value)
}

func (c cons) String(name string, value string) {
	c.print(name, value)
}

func (c cons) Any(name string, value interface{}) {
	c.print(name, value)
}
