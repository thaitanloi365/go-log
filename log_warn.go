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

func (l *Logger) WarnWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, "", values...).withTag(tag)
}

func (l *Logger) WarnfWithTag(tag string, format string, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeInterface, format, values...).withTag(tag)
}

func (l *Logger) WarnJSONWithTag(tag string, values ...interface{}) {
	l.queue <- l.buildlog(Warn, l.fileWithLineNum(), valueTypeJSON, "", values...).withTag(tag)
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
