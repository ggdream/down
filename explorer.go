package main

import "io"

type reader struct {
	io.Reader

	Total   int64
	Current int64

	CallBack func(c *context)
}

func (r *reader) Read(data []byte) (int, error) {
	n, err := r.Reader.Read(data)
	if err != nil {
		return 0, err
	}
	r.Current += int64(n)

	c := &context{
		Total:   r.Total,
		Current: r.Current,
	}
	r.CallBack(c)

	return n, nil
}
