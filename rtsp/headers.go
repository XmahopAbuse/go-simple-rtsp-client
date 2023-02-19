package rtsp

type RTSPHeaders struct {
	CSeq string
	UserAgent string
	Accept string
	ContentType string
	Authorization string

}

func (h *RTSPHeaders) SetHeaders(req *RTSPRequest) {
	if h.CSeq != "" {
		req.Headers["CSeq"] = h.CSeq
	}
	if h.UserAgent != "" {
		req.Headers["User-Agent"] = h.UserAgent
	}
	if h.Accept != "" {
		req.Headers["Accept"] = h.Accept
	}
	if h.ContentType != "" {
		req.Headers["Content-Type"] = h.ContentType
	}
}