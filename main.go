package main

import (
	"log"
	"rtspClient/rtsp"
)

func main() {
	address := "rtsp://admin:a0w5fj7o09@172.16.41.31:554/Streaming/Channels/101"
	client, err := rtsp.NewClient(address)
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	client.Describe()
}