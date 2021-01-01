package main

import "time"

type (
	mapping = map[string]string
	meta    struct {
		TimeOut  time.Duration
		SavePath string
		CallBack func(c *context)
	}
)
