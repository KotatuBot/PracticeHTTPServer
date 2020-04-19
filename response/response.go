package response

import (
	"fmt"
	"io/ioutil"
	"os"

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

func CreateHeader(requesthead header.HttpHeader, status string, filename string, Body string) string {
	var Header string

	headers := header.ResponseHeader{}
	headers.Status_Code = "HTTP1.1" + " " + status + " " + status
	headers.Access_Control_Allow_Origin = "*"
	headers.Connection = "keep-alive"
	//headers.Content_Encoding = SetContentEncoding(requesthead.Accept_Encodeing)
	headers.Content_Type.Media_Type, headers.Content_Type.Charset = SetContentType(filename)
	headers.Date = SetDatetime()
	headers.Etag = SetEtag(Body)
	headers.Keep_Alive.Timeout, headers.Keep_Alive.Max = SetKeepAlive()
	headers.Server = "Scrach"
	headers.X_Frame_Options = SetXFrameOptions("same")
	Cookie_Data := SetCookie()
	headers.Set_Cookie = Cookie_Data

	Header = fmt.Sprintf("%s\r\nAccess-Control-Allow-Origin: %s\r\nConnection: %s\r\nContent-Type: %scharset=%s\r\nConnection: %s\r\nDate: %s\r\nEtag: %s\r\nKeep-Alive: %s,%s\r\nServer: %s\r\nX-Frame-Options: %s\r\nSet-Cookie: %s=%s;Expires=%s;%s;SameSite=%s\r\n", headers.Status_Code, headers.Access_Control_Allow_Origin, headers.Connection, headers.Content_Type.Media_Type, headers.Content_Type.Charset, headers.Connection, headers.Date, headers.Etag, headers.Keep_Alive.Timeout, headers.Keep_Alive.Max, headers.Server, headers.X_Frame_Options, headers.Set_Cookie.Name, headers.Set_Cookie.Value, headers.Set_Cookie.Expires, headers.Set_Cookie.HttpOnly, headers.Set_Cookie.SameSiteCookie)
	Header += "\r\n\r\n"
	return Header
}

func ResponseMessage(requesthead header.HttpHeader) []byte {

	var Header, Body, Message string
	var response []byte
	var filename string
	var status_code string

	file_path := requesthead.Url
	// Request Parse
	filename, status_code = router.RouterPath(file_path)

	//Body
	Body = FileRead(filename)
	// Head
	Header = CreateHeader(requesthead, status_code, filename, Body)
	Header += " \r\n\r\n"

	Message = Header + Body
	fmt.Println("Response")
	fmt.Println(Message)
	response = []byte(Header)
	body_bytes := []byte(Body)
	response = append(response, body_bytes...)

	return response
}
