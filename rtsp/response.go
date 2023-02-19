package rtsp

type RTSPResponse struct {
	StatusCode int
	Headers map[string]string
	Body string
}

func ParseRTSPResponse() {

}