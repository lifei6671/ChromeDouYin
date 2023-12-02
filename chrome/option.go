package chrome

import (
	"time"
)

type options struct {
	ServiceURL string `json:"service_url"`
	Cookies    []Cookie
	Timeout    time.Duration
}

type Cookie struct {
	Name   string `json:"name"`
	Cookie string `json:"cookie"`
	Domain string `json:"domain"`
}

type Option func(*options)

func WithCookies(c ...Cookie) Option {
	return func(o *options) {
		o.Cookies = append(o.Cookies, c...)
	}
}

func WithServiceURL(c string) Option {
	return func(o *options) {
		o.ServiceURL = c
	}
}

func WithTimeout(d time.Duration) Option {
	return func(o *options) {
		o.Timeout = d
	}
}
