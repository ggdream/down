package main

import (
	"io"
	"net/http"
	"os"
)

func fetcher(url string, headers mapping, info meta) error {
	c := client{
		&http.Client{},
	}

	res, err := c.Get(url, headers)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rer := &reader{
		Reader:   res.Body,
		Total:    res.ContentLength,
		CallBack: info.CallBack,
	}
	file, err1 := os.Create(info.SavePath)
	if err1 != nil {
		return err1
	}
	defer file.Close()

	_, err2 := io.Copy(file, rer)
	return err2
}
