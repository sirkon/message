package testing_pkg

import (
	"os"
	"github.com/sirkon/message"
)

func ExampleLocation() {
	os.Stderr = os.Stdout
	message.SetDebug(true)
	message.UseColor(false)
	message.Info("message")
	// Output: /home/emacs/go/src/github.com/sirkon/message/testing_pkg/print_test.go:12 message
}
