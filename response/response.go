package response

import (
	"io/ioutil"
	"os"
)

func FileRead(filename string) string {

	var filepath string
	var filecontent string

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filepath = pwd + "/File/" + filename
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	filecontent = string(bytes)
	return filecontent

}

func ResponseMessage(file_path string) []byte {

	var Header, Body, Message string
	var response []byte
	var filename string

	Header = ""
	Header += "HTTP1.1"
	Header += " 200"
	Header += " 200\r\n"
	Header += " \r\n\r\n"

	if file_path == "/" {

		filename = "test.html"
	} else {

		filename = "test.html"
	}
	Body = FileRead(filename)

	Message = Header + Body

	response = []byte(Message)

	return response
}
