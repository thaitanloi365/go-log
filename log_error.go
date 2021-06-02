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

func (l *Logger) ErrorWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, "", values...).withTag(tag)
}

func (l *Logger) ErrorfWithTag(tag string, format string, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeInterface, format, values...).withTag(tag)
}

func (l *Logger) ErrorJSONWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Error, l.fileWithLineNum(), valueTypeJSON, "", values...).withTag(tag)
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
