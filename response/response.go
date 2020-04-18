package response

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"../header"
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
func SetStatusCode(statuscode string) string {
	var statusmessage string
	/*
		var status_condition string

		switch statuscode {

		case "200":
			status_condition = "OK"
		case "404":
			status_condition = "Not Found"
		default:
			status_condition = "NG"
		}
	*/

	statusmessage = "HTTP1.1 " + statuscode + " " + statuscode

	return statusmessage
}

func SetContentEncoding(requestencoding string) string {

	var content_encoding string

	available := strings.Split(requestencoding, ",")
	content_encoding = available[0]

	return content_encoding

}

func SetContentType(filename string) (string, string) {
	var content_type string
	var charset string
	var length int

	file_types := strings.Split(filename, ".")
	length = len(file_types) - 1

	content_type = file_types[length]

	switch content_type {

	case "html":
		content_type = "text/html;"
	case "ico":
		content_type = "image/x-icon;"
	default:
		content_type = "text/html;"

	}
	charset = "UTF-8"

	return content_type, charset

}

func SetDatetime() string {
	now_times := time.Now()
	weekday := now_times.Weekday()
	day := now_times.Day()
	month := now_times.Month()
	Year := now_times.Year()
	hour := now_times.Hour()
	minute := now_times.Minute()
	second := now_times.Second()
	TimCode := "JST"

	dates := fmt.Sprintf("%s, %d %d %d %d:%d:%d %s", weekday, day, month, Year, hour, minute, second, TimCode)

	return dates
}

func CreateHeader(requesthead header.HttpHeader, status string, filename string) string {
	var Header string

	headers := header.ResponseHeader{}
	headers.Status_Code = "HTTP1.1" + " " + status + " " + status
	headers.Access_Control_Allow_Origin = "*"
	headers.Connection = "keep-alive"
	//headers.Content_Encoding = SetContentEncoding(requesthead.Accept_Encodeing)
	headers.Content_Type.Media_Type, headers.Content_Type.Charset = SetContentType(filename)
	headers.Date = SetDatetime()

	Header = fmt.Sprintf("%s\r\nAccess-Control-Allow-Origin: %s\r\nContent-Type: %scharset=%s\r\nConnection: %s\r\nDate: %s\r\n", headers.Status_Code, headers.Access_Control_Allow_Origin, headers.Content_Type.Media_Type, headers.Content_Type.Charset, headers.Connection, headers.Date)
	Header += "\r\n\r\n"
	return Header
}

func ResponseMessage(requesthead header.HttpHeader) []byte {

	var Header, Body, Message string
	var response []byte
	var filename string
	var status_code string

	file_path := requesthead.Url

	filename, status_code = router.RouterPath(file_path)
	Header = CreateHeader(requesthead, status_code, filename)
	Header += " \r\n\r\n"

	Body = FileRead(filename)

	Message = Header + Body
	fmt.Println("Response")
	fmt.Println(Message)

	response = []byte(Message)

	return response
}
