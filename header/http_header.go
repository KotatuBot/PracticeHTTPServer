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

type KeepAlive struct {
	Timeout string
	Max     string
}

type Cookie struct {
	Name           string
	Value          string
	Expires        string
	Domain         string
	Path           string
	HttpOnly       string
	Secure         string
	SameSiteCookie string
}

type ResponseHeader struct {
	Status_Code                 string
	Access_Control_Allow_Origin string
	Connection                  string
	Content_Encoding            string
	Content_Type                ContentTypes
	Date                        string
	Etag                        string
	Keep_Alive                  KeepAlive
	Last_Modified               string
	Server                      string
	Set_Cookie                  Cookie
	Transfer_Encoding           string
	Cache_Control               string
	X_Frame_Options             string
}
