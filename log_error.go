package logger

func (l *Logger) Error(values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

func (l *Logger) Errorf(format string, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

func (l *Logger) ErrorJSON(values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

func (l *Logger) ErrorWithOption(option *Option, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, "", values...).withOption(option)
}

func (l *Logger) ErrorfWithOption(option *Option, format string, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, format, values...).withOption(option)
}

func (l *Logger) ErrorJSONWithOption(option *Option, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeJSON, "", values...).withOption(option)
}

func (l *Logger) ErrorWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

func (l *Logger) ErrorfWithRequestInfo(reqInfo *RequestInfo, format string, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, format, values...).withRequestInfo(reqInfo)
}

func (l *Logger) ErrorJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}
