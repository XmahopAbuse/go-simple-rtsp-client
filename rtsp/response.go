package rtsp

import "fmt"

type RTSPResponse struct {
	StatusCode int
	Headers map[string]string
	Body string
}


func ParseRTSPResponse(resp []byte) {
	fmt.Println(string(resp))
}