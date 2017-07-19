# message
Not everyone needs configured logging.

```bash
go get github.com/glossina/message
```

This library is primarily for command line short lived utilities where the time output is not needed. Only message itself.

What is supported:

* Leveled output distincted with colors

    | level | color | methods |
    |:-----:|:-----:|:-------:|
    |critical|bold red|Critical, Criticalf, Fatal, Fatalf|
    |error|red|Error, Errorf|
    |warning|yellow|Warning, Warningf|
    |notice|green|Notice, Noticef|
    |info|reset|Info, Infof|
    |debug|cyan|Debug, Debugf|

* Code location output - DEBUG environment variable must be set into not empty value

### Usage example
```go
package main

import "github.com/glossina/message"

func main() {
	message.Notice("Hello world!")
}
```
