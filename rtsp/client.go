package rtsp

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/url"
	"strings"
)

type Client struct {
	Address string
	Username string
	Password string
	Stream string
	Scheme string
	Path string
	conn net.Conn
}


func NewClient(rawUrl string) (*Client, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	password, _  := u.User.Password()

	return &Client{Address: u.Host,
			       Username: u.User.Username(),
				   Password: password,
				   Scheme: u.Scheme,
				   Path: u.Path,
	}, nil

}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

func (c *Client) Options() error {

	headers := &RTSPHeaders{
		CSeq: "1",
		UserAgent: "test",
		Accept: "application/sdp",
	}
	req := &RTSPRequest{
		Headers: make(map[string]string),
		Method: Options,
		Url: fmt.Sprintf("%s://%s%s", c.Scheme, c.Address, c.Path),
	}

	headers.SetHeaders(req)

	payload := req.Marshall()

	buff := make([]byte, 1024)
	c.conn.Write([]byte(payload))
	_, err := c.conn.Read(buff)
	if err != nil {
		return err
	}
	fmt.Println(string(buff))
	return nil
}

func (c *Client) Describe() (err error) {
	c.Options()

	headers := &RTSPHeaders{
		CSeq:      "2",
		UserAgent: "test",
		Accept:    "application/sdp",
		Authorization: fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.Username, c.Password)))),
	}

	req := &RTSPRequest{
		Headers: make(map[string]string),
		Method:  Describe,
		Url:     fmt.Sprintf("%s://%s%s", c.Scheme, c.Address, c.Path),
	}

	headers.SetHeaders(req)

	payload := req.Marshall()

	buff := make([]byte, 1024)
	_, err = c.conn.Write([]byte(payload))

	n, err := c.conn.Read(buff)

	// Check if the response is 401 Unauthorized
	if strings.Contains(string(buff[:n]), "RTSP/1.0 401 Unauthorized") {
		if strings.Contains(string(buff[:n]), "WWW-Authenticate: Digest") {
			authHeader := strings.Split(string(buff[:n]), "WWW-Authenticate: Digest ")[1]
			fmt.Println(authHeader, "header")
		}
	}

	return  nil
}