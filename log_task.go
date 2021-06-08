package logger

import (
	"fmt"
	"strings"
)

// LogLevel log level
type valueType int

const (
	valueTypeInterface valueType = iota + 1
	valueTypeJSON
	valueTypeCustom
)

type logTask struct {
	logger      *Logger
	logLevel    LogLevel
	time        string
	format      string
	values      []interface{}
	caller      string
	valueType   valueType
	requestInfo *RequestInfo
	option      *Option
}

func (task *logTask) withRequestInfo(requestInfo *RequestInfo) *logTask {
	task.requestInfo = requestInfo
	return task
}

func (task *logTask) withOption(option *Option) *logTask {
	task.option = option
	return task
}

func (task *logTask) formatRequestInfo() string {
	if task.requestInfo == nil {
		return ""
	}
	var info = task.requestInfo

	var extras = []string{}
	if info.UserID != "" {
		extras = append(extras, info.UserID)
	}

	if info.RefErrorID != "" {
		extras = append(extras, info.RefErrorID)
	}

	if len(extras) > 0 {
		return fmt.Sprintf("%s [%v] %d %s %s", info.ReqID, strings.Join(extras, "::"), info.Status, info.Method, info.URI)
	}

	return fmt.Sprintf("%s %d %s %s", info.ReqID, info.Status, info.Method, info.URI)
}

func (task *logTask) formatOption() string {
	if task.option == nil {
		return ""
	}
	return task.option.Format()
}
