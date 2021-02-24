package logger

// Notifier interface
type Notifier interface {
	Send(title, body string) error
}
