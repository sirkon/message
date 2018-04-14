package message

import (
	"fmt"
	"strconv"
	"strings"
)

var bslash033Lsbrck = "\033["

// ColorInfo ...
type ColorInfo struct {
	rest string
	Text string
	Code int 
}

// Extract ...
func (p *ColorInfo) Extract(line string) (bool, error) {
	p.rest = line
	var err error
	var pos int
	var tmp string
	var tmpInt int64

	// Take until "\033[" as Text(string)
	pos = strings.Index(p.rest, bslash033Lsbrck)
	if pos >= 0 {
		p.Text = p.rest[:pos]
		p.rest = p.rest[pos+len(bslash033Lsbrck):]
	} else {
		return false, nil
	}

	// Take until 'm' as Code(int)
	pos = strings.IndexByte(p.rest, 'm')
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}
	if tmpInt, err = strconv.ParseInt(tmp, 10, 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Code = int(tmpInt)

	return true, nil
}
