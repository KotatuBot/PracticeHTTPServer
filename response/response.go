package response

func ResponseMessage() []byte {

	var Header, Body, Message string
	var response []byte

	Header = ""
	Header += "HTTP1.1"
	Header += " 200"
	Header += " 200\r\n"
	Header += " \r\n\r\n"
	Body = "<h1>Hello Kotatu</h1>\r\n<p>Net Work</p>\r\n"

	Message = Header + Body

	response = []byte(Message)

	return response
}
