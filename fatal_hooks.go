package message

var fatalHooks []func()

// AddFatalHook adds a function to execute before
// the os.Exit(1) in Fatal family of functions.
func AddFatalHook(f func()) {
	fatalHooks = append(fatalHooks, f)
}

// RemoveFatalHooks clears all fatal hooks
func RemoveFatalHooks() {
	fatalHooks = fatalHooks[:0]
}

func execHooks() {
	for _, hook := range fatalHooks {
		hook()
	}
}
