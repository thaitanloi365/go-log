package logger

func (l *Logger) Info(values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

func (l *Logger) Infof(format string, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

func (l *Logger) InfoJSON(values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

func (l *Logger) InfoWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, "", values...).withTag(tag)
}

func (l *Logger) InfofWithTag(tag string, format string, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, format, values...).withTag(tag)
}

func (l *Logger) InfoJSONWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeJSON, "", values...).withTag(tag)
}

func (l *Logger) InfoWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

func (l *Logger) InfofWithRequestInfo(reqInfo *RequestInfo, format string, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeInterface, format, values...).withRequestInfo(reqInfo)
}

func (l *Logger) InfoJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Info, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}
