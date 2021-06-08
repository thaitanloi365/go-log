package logger

import "time"

type RequestInfo struct {
	ReqID        string
	Status       int
	Method       string
	URI          string
	UserID       string
	RefErrorID   string
	ResponseTime time.Time

	option *Option
}

func (r *RequestInfo) WithOption(option *Option) *RequestInfo {
	r.option = option
	return r
}
