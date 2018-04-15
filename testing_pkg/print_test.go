package testing_pkg

import (
	"github.com/sirkon/message"
	"bytes"
	"testing"
	"os"
	"os/user"
	"path/filepath"
	"fmt"
)

func TestPrinter(t *testing.T) {
	buf := &bytes.Buffer{}
	message.SetDebug(true)
	message.UseColor(false)
	message.SetDest(buf)
	message.Info("message")
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		usr, err := user.Current()
		if err != nil {
			t.Fatalf("cannot get current user: %s", err)
		}
		gopath = filepath.Join( usr.HomeDir, "go")
	}
	filePath := filepath.Join(gopath, "src", "github.com", "sirkon", "message", "testing_pkg", "print_test.go")
	// output must be $GOPATH/src/github.com/sirkon/message/testing_pkg/print_test.go:12 message with new line
	expected := fmt.Sprintf("%s:%d message\n", filePath, 18)
	if expected != buf.String() {
		t.Fatalf("output must be %s, got %s", expected, buf.String())
	}

}
