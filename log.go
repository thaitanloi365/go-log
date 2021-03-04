package logger

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// LogLevel log level
type LogLevel int

// All levels
const (
	Debug LogLevel = iota + 1
	Warn
	Info
	Error
)

// Logger instance
type Logger struct {
	context    context.Context
	cancelFunc context.CancelFunc

	mutex sync.RWMutex

	queue chan *logTask

	writer     Writer
	fileWriter Writer

	debugStr      string
	debugColorStr string

	infoStr      string
	infoColorStr string

	warnStr      string
	warnColorStr string

	errStr      string
	errColorStr string

	// config
	colorful     bool
	timeLocation *time.Location
	dateFormat   string
	prefix       string

	writeFileExceptLevels []LogLevel

	notifier Notifier
}

// Config config
type Config struct {
	BufferedSize int
	Colorful     bool
	TimeLocation *time.Location
	DateFormat   string
	Prefix       string

	Writer                Writer
	WriteFileExceptLevels []LogLevel

	Notifier Notifier
}

// New init logger
func New(config *Config) *Logger {
	ctx, cancelFunc := context.WithCancel(context.Background())
	var logger = &Logger{
		writer:        log.New(os.Stdout, "\r\n", 0),
		fileWriter:    log.New(ioutil.Discard, "", 0),
		mutex:         sync.RWMutex{},
		context:       ctx,
		cancelFunc:    cancelFunc,
		queue:         make(chan *logTask, 10),
		debugStr:      debugStr,
		debugColorStr: debugColorStr,
		infoStr:       infoStr,
		infoColorStr:  infoColorStr,
		warnStr:       warnStr,
		warnColorStr:  warnColorStr,
		errStr:        errStr,
		errColorStr:   errColorStr,
		timeLocation:  time.Local,
		colorful:      true,
		dateFormat:    time.RFC3339,
		writeFileExceptLevels: []LogLevel{
			Debug,
		},
	}

	if config != nil {
		if config.BufferedSize > 0 {
			logger.queue = make(chan *logTask, config.BufferedSize)
		}

		if config.DateFormat != "" {
			logger.dateFormat = config.DateFormat
		}

		if config.TimeLocation == nil {
			logger.timeLocation = config.TimeLocation
		}

		if config.Writer != nil {
			logger.fileWriter = config.Writer
		}

		if config.WriteFileExceptLevels != nil {
			logger.writeFileExceptLevels = config.WriteFileExceptLevels
		}

		logger.notifier = config.Notifier

	}

	logger.run()

	return logger
}

func (l *Logger) fileWithLineNum() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)

		if ok {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}

func (l *Logger) buildlog(logtype LogLevel, caller string, valueType valueType, format string, values ...interface{}) *logTask {
	var now = time.Now()
	if l.timeLocation != nil {
		now = now.In(l.timeLocation)
	}
	return &logTask{
		logger:    l,
		logLevel:  logtype,
		time:      now.Format(l.dateFormat),
		format:    format,
		values:    values,
		caller:    caller,
		valueType: valueType,
	}

}

// Print print
func (l *Logger) Print(values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeCustom, "", values...)
}

// Printf printf
func (l *Logger) Printf(format string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeCustom, format, values...)
}

// Debug debug
func (l *Logger) Debug(values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

// Debugf debug with format
func (l *Logger) Debugf(format string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

// DebugJSON print pretty json
func (l *Logger) DebugJSON(values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

// Info info
func (l *Logger) Info(values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

// Infof info with format
func (l *Logger) Infof(format string, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

// InfoJSON print pretty json
func (l *Logger) InfoJSON(values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

// Warn warn
func (l *Logger) Warn(values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

// Warnf info with format
func (l *Logger) Warnf(format string, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

// WarnJSON print pretty json
func (l *Logger) WarnJSON(values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

// Error error
func (l *Logger) Error(values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

// Errorf error with format
func (l *Logger) Errorf(format string, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

// ErrorJSON print pretty json
func (l *Logger) ErrorJSON(values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

// DebugWithRequestInfo debug with request info
func (l *Logger) DebugWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

// ErrorWithRequestInfo debug with request info
func (l *Logger) ErrorWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

// InfoWithRequestInfo debug with request info
func (l *Logger) InfoWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

// WarnWithRequestInfo debug with request info
func (l *Logger) WarnWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

// DebugJSONWithRequestInfo debug with request info
func (l *Logger) DebugJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}

// ErrorJSONWithRequestInfo debug with request info
func (l *Logger) ErrorJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}

// InfoJSONWithRequestInfo debug with request info
func (l *Logger) InfoJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}

// WarnJSONWithRequestInfo debug with request info
func (l *Logger) WarnJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}

func (l *Logger) run() {
	go l.cleanup()

	go func(ctx context.Context, queue chan *logTask) {
		for {
			select {
			case <-ctx.Done():
				return

			case data := <-queue:
				var format = l.infoStr
				var formatColor = l.infoColorStr
				var extraFormat = data.format
				var extraPrettyFormat = data.format
				switch data.logLevel {
				case Debug:
					format = l.debugStr
					formatColor = l.debugColorStr
				case Error:
					format = l.errStr
					formatColor = l.errColorStr
				case Warn:
					format = l.warnStr
					formatColor = l.warnColorStr
				}

				var separator = " "
				switch data.valueType {
				case valueTypeJSON:
					separator = "\n"
				}

				var fullFormatColor = formatColor
				var fullFormat = format

				if data.requestInfo != nil {
					fullFormatColor = formatColor + data.formatRequestInfo() + "\n"
					fullFormat = fullFormat + data.formatRequestInfo() + " "
				}

				if extraPrettyFormat == "" {
					for i := 0; i < len(data.values); i++ {
						extraPrettyFormat = "%v" + separator + extraPrettyFormat
					}
				}

				if extraFormat == "" {
					for i := 0; i < len(data.values); i++ {
						extraFormat = "%v" + " " + extraFormat
					}
				}

				fullFormatColor = fullFormatColor + extraPrettyFormat
				fullFormat = fullFormat + extraFormat

				switch data.valueType {
				case valueTypeCustom:
					l.writer.Printf(fullFormatColor, append([]interface{}{data.time, data.caller}, data.values...)...)

					if l.isIgnoreWriteFile(data.logLevel) == false {
						l.fileWriter.Printf(fullFormat, append([]interface{}{data.time, data.caller}, data.values...)...)
					}

					if l.notifier != nil {
						var titleFormat = format
						if data.requestInfo != nil {
							titleFormat = data.formatRequestInfo() + "\n" + titleFormat
						}

						l.notifier.Send(fmt.Sprintf(titleFormat, data.time), fmt.Sprintf(data.format, data.values...))
					}
				case valueTypeJSON:
					var prettyValues = []interface{}{}
					var values = []interface{}{}
					for _, value := range data.values {
						switch v := value.(type) {
						default:
							values = append(values, ToJSONString(value))
							prettyValues = append(prettyValues, ToPrettyJSONString(value))
						case uint64, string, int, int64, int32, bool, float32, float64:
							values = append(values, v)
							prettyValues = append(prettyValues, v)
						}

					}
					l.writer.Printf(fullFormatColor, append([]interface{}{data.time, data.caller}, prettyValues...)...)

					if l.isIgnoreWriteFile(data.logLevel) == false {
						l.fileWriter.Printf(fullFormat, append([]interface{}{data.time, data.caller}, values...)...)
					}

					if l.notifier != nil {
						var titleFormat = format
						if data.requestInfo != nil {
							titleFormat = data.formatRequestInfo() + "\n" + titleFormat
						}

						l.notifier.Send(fmt.Sprintf(titleFormat, data.time, data.caller), fmt.Sprintf(extraFormat, prettyValues...))
					}
				default:
					l.writer.Printf(fullFormatColor, append([]interface{}{data.time, data.caller}, data.values...)...)

					if l.isIgnoreWriteFile(data.logLevel) == false {
						l.fileWriter.Printf(fullFormat, append([]interface{}{data.time, data.caller}, data.values...)...)
					}

					if l.notifier != nil {
						var titleFormat = format
						if data.requestInfo != nil {
							titleFormat = data.formatRequestInfo() + "\n" + titleFormat
						}
						l.notifier.Send(fmt.Sprintf(titleFormat, data.time, data.caller), fmt.Sprintf(extraFormat, data.values...))
					}

				}

				break
			}
		}
	}(l.context, l.queue)
}

func (l *Logger) isIgnoreWriteFile(level LogLevel) bool {
	for _, lv := range l.writeFileExceptLevels {
		if lv == level {
			return true
		}
	}

	return false

}

func (l *Logger) cleanup() {
	<-l.context.Done()

	// Lock the destinations
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// Cleanup the destinations
	close(l.queue)

}
