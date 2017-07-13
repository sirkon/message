/*
* THE FILE WAS GENERATED WITH /home/emacs/Sources/indep/bin/logparsergen --source=src/github.com/DenisCheremisov/ch-insert/message/p.script --package=message
* DO NOT TOUCH IT!
 */
package message

import (
	"bytes"
)

var const0 = []byte("\033[")

type colorExtractor struct {
	rest []byte
	Head []byte
	Code []byte
	Rest []byte
}

func (p *colorExtractor) Parse(line []byte) (bool, error) {
	p.rest = line
	var pos int
	if pos = bytes.Index(p.rest, const0); pos < 0 {
		return false, nil
	}
	p.Head = p.rest[:pos]
	p.rest = p.rest[pos+len(const0):]
	if len(p.rest) < 4 {
		return false, nil
	}
	if pos = bytes.IndexByte(p.rest[:4], 'm'); pos < 0 {
		return false, nil
	}
	p.Code = p.rest[:pos]
	p.rest = p.rest[pos+1:]
	p.Rest = p.rest
	p.rest = p.rest[len(p.rest):]
	return true, nil
}
