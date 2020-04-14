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
