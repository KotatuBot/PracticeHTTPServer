package response

import (
	"io/ioutil"
	"os"

	"../router"
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

func CreateHeader(status string) string {
	var Header string
	Header = ""
	Header += "HTTP1.1"
	Header += " " + status
	Header += " " + status + "\r\n"
	Header += " \r\n\r\n"
	return Header

}

func ResponseMessage(file_path string) []byte {

	var Header, Body, Message string
	var response []byte
	var filename string
	var status_code string

	Header = ""
	Header += "HTTP1.1"
	Header += " 200"
	Header += " 200\r\n"
	Header += " \r\n\r\n"

	filename, status_code = router.RouterPath(file_path)
	Header = CreateHeader(status_code)

	Body = FileRead(filename)

	Message = Header + Body

	response = []byte(Message)

	return response
}
