package main

import (
	"time"
)

func run(url, savePath string, timeOut time.Duration, callback func(c *context)) error {
	info := meta{
		TimeOut:  timeOut,
		SavePath: savePath,
		CallBack: callback,
	}

	return fetcher(url, headers, info)
}
