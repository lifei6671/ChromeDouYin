package chrome

import (
	"time"
)

type options struct {
	ServiceURL string `json:"service_url"`
	Cookie     string `json:"cookie"`
	Timeout    time.Duration
}

type Option func(*options)

func WithCookie(c string) Option {
	return func(o *options) {
		o.Cookie = c
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
