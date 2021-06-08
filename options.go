package logger

import "fmt"

type Option struct {
	ID          string
	Tag         string
	Description string
}

func NewOption(id ...string) *Option {
	var opt = &Option{}
	if len(id) > 0 {
		opt.ID = id[0]
	}
	return opt
}

func (opt *Option) WithTag(tag string) *Option {
	opt.Tag = tag
	return opt
}

func (opt *Option) WithID(id string) *Option {
	opt.ID = id
	return opt
}

func (opt *Option) WithDescription(description string) *Option {
	opt.Description = description
	return opt
}

func (opt *Option) Format() string {
	return fmt.Sprintf("[%s::%s] %s", opt.ID, opt.Tag, opt.Description)
}
