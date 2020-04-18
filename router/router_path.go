package router

import (
	"strings"
)

func RouterPath(recv_path string) (string, string) {

	var filename string
	var routename string
	var length int
	var status_code string

	status_code = "200"

	file_array := strings.Split(recv_path, "/")
	length = len(file_array) - 1

	if len(file_array[length]) == 0 {
		routename = "/"

	} else {

		routename = file_array[length]
	}

	switch routename {

	case "/":
		filename = "index.html"
	case "index.html":
		filename = "index.html"
	default:
		filename = "notfound.html"
		status_code = "404"

	}

	return filename, status_code

}
