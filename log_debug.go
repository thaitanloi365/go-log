package logger

func (l *Logger) Debug(values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

func (l *Logger) Debugf(format string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

func (l *Logger) DebugJSON(values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

func (l *Logger) DebugWithOption(option *Option, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, "", values...).withOption(option)
}

func (l *Logger) DebugfWithOption(option *Option, format string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, format, values...).withOption(option)
}

func (l *Logger) DebugJSONWithOption(option *Option, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeJSON, "", values...).withOption(option)
}

func (l *Logger) DebugWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

func (l *Logger) DebugfWithRequestInfo(reqInfo *RequestInfo, format string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, format, values...).withRequestInfo(reqInfo)
}

func (l *Logger) DebugJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}
