package logger

func (l *Logger) Warn(values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, "", values...)
}

func (l *Logger) Warnf(format string, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, format, values...)
}

func (l *Logger) WarnJSON(values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeJSON, "", values...)
}

func (l *Logger) WarnWithTag(option *Option, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, "", values...).withOption(option)
}

func (l *Logger) WarnfWithTag(option *Option, format string, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, format, values...).withOption(option)
}

func (l *Logger) WarnJSONWithTag(option *Option, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeJSON, "", values...).withOption(option)
}

func (l *Logger) WarnWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, "", values...).withRequestInfo(reqInfo)
}

func (l *Logger) WarnfWithRequestInfo(reqInfo *RequestInfo, format string, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, format, values...).withRequestInfo(reqInfo)
}

func (l *Logger) WarnJSONWithRequestInfo(reqInfo *RequestInfo, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeJSON, "", values...).withRequestInfo(reqInfo)
}
