package main

import "fmt"

const (
	name    = "down"
	version = "v0.0.1"
)

var (
	ua = fmt.Sprintf("%s/%s", name, version)
)
