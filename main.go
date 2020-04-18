package main

import (
	"io"
	"log"
	"net"
	"os"

	"./header"
	"./parse"
	"./response"
)

func StartServer() {

	ln, err := net.Listen("tcp", "localhost:8080")
	defer ln.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {

		heads := header.HttpHeader{}

		conn, err := ln.Accept()
		if err != nil {

			log.Println(err)
		}

		// Client read Data
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		recv_data := parse.RecvData(buf, n)
		heads = parse.UrlParse(recv_data)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}

		message := response.ResponseMessage(heads.Url)
		conn.Write(message)
		conn.Close()

	}

}

func main() {
	// ForLooping
	for {
		StartServer()
	}

}
