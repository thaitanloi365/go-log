package logger

// Writer interface
type Writer interface {
	Printf(string, ...interface{})
	Print(...interface{})
}
