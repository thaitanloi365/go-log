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
	Tag          string
}

func (r *RequestInfo) WithTag(tag string) *RequestInfo {
	r.Tag = tag
	return r
}
