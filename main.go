package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	callback := func(c *context) {
		fmt.Printf("\b\b\b\b\b%2.2f", float64(c.Current*10000/c.Total)/100)
	}

	s := strings.Split(strings.Split(os.Args[1], "?")[0], "/")
	if err := run(os.Args[1], s[len(s)-1], 3*time.Second, callback); err != nil {
		panic(err.Error())
	}
}
