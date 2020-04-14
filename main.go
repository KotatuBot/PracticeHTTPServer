package main

import (
	"./parse"
	"./response"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	defer ln.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {

		conn, err := ln.Accept()
		if err != nil {

			log.Println(err)
		}

		// Client read Data
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		recv_data := parse.RecvData(buf, n)
		parse.UrlParse(recv_data)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}

		message := response.ResponseMessage()
		conn.Write(message)
		conn.Close()

	}

}
