package header

type HttpHeader struct {
	Method                    string
	Url                       string
	Version                   string
	Host                      string
	Useragent                 string
	Accept                    string
	Accept_Language           string
	Accept_Encodeing          string
	Connection                string
	Cookie                    string
	Upgrade_Insecure_Requests string
	Cache_Control             string
	If_None_Match             string
}
type ContentTypes struct {
	Media_Type string
	Charset    string
}

type ResponseHeader struct {
	Status_Code                 string
	Access_Control_Allow_Origin string
	Connection                  string
	Content_Encoding            string
	Content_Type                ContentTypes
	Date                        string
	Etag                        string
	Keep_Alive                  string
	Last_Modified               string
	Server                      string
	Set_Cookie                  string
	Transfer_Encoding           string
	Vary                        string
	X_Backend_Server            string
	X_Cache_Info                string
	X_kuma_Revision             string
	X_frame_option              string
}
