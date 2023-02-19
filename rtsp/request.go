package rtsp

import (
	"fmt"
	"strings"
)

type Method string

const RTSPVersion = "RTSP/1.0"

const (
Options Method = "OPTIONS"
Describe Method = "DESCRIBE"
)



type RTSPRequest struct {
	Method Method
	Url string
	Headers map[string]string
}

func (r *RTSPRequest) Marshall() string{
	var b strings.Builder

	fmt.Fprintf(&b, "%s %s %s\r\n", r.Method, r.Url, RTSPVersion)
	for k, v := range r.Headers {
		fmt.Fprintf(&b, "%s: %s\r\n", k, v)
	}

	fmt.Fprintf(&b, "\r\n")

	return b.String()
}