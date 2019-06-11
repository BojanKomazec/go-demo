package onerr

// PanicOnError function implements a common error handler.
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
