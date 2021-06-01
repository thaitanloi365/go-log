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

func (l *Logger) DebugWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, "", values...).withTag(tag)
}

func (l *Logger) DebugfWithTag(tag string, format string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeInterface, format, values...).withTag(tag)
}

func (l *Logger) DebugJSONWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Debug, l.fileWithLineNum(), valueTypeJSON, "", values...).withTag(tag)
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
