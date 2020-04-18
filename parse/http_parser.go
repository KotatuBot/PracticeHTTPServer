package parse

import (
	"fmt"
	"strings"

	"../header"
)

// Recv Message -> From Byte to String
func RecvData(buf []byte, position int) string {
	var recv_message string
	recv_message = string(buf[:position])
	return recv_message
}

func HeadParse(headers *header.HttpHeader, recv string) {

	var header []string

	//Parse recv data
	header = strings.Split(recv, ": ")

	switch header[0] {

	case "Host":
		headers.Host = header[1]
	case "User-Agent":
		headers.Useragent = header[1]

	case "Accept":
		headers.Accept = header[1]

	case "Accept-Language":
		headers.Accept_Language = header[1]

	case "Accept-Encoding":
		headers.Accept_Encodeing = header[1]

	case "Connection":
		headers.Connection = header[1]

	case "cookie":
		headers.Cookie = header[1]

	case "Upgrade-Insecure-Requests":
		headers.Upgrade_Insecure_Requests = header[1]

	case "Cache-Control":
		headers.Cache_Control = header[1]

	case "If-None-Match":
		headers.If_None_Match = header[1]

	default:
		fmt.Println("Not Found")

	}

}

func UrlParse(recv_message string) header.HttpHeader {

	var recv_array []string

	headers := header.HttpHeader{}

	recv_array = strings.Split(recv_message, "\r\n")

	for number, recv := range recv_array {
		// Method
		if number == 0 {
			method_data := strings.Split(recv, " ")
			if len(method_data) == 3 {
				headers.Method = method_data[0]
				headers.Url = method_data[1]
				headers.Version = method_data[2]
			}
		} else {
			//Head Parse
			HeadParse(&headers, recv)

		}

	}

	fmt.Println("-----------")
	fmt.Println(headers)
	return headers

}
