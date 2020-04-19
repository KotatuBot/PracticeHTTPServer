package response

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

func SetStatusCode(statuscode string) string {
	var statusmessage string

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

func SetEtag(body string) string {
	var etag string

	etag = fmt.Sprintf("\"%x\"", md5.Sum([]byte(body)))
	return etag
}

func SetKeepAlive() (string, string) {
	var timeout string
	var maxconnection string

	timeout = "timeout=5"
	maxconnection = "max=100"
	return timeout, maxconnection

}

func SetCacheControl() string {
	return "no-store"

}
