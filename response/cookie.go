package response

import (
	"fmt"
	"time"

	"../header"
)

func SetCookie() header.Cookie {

	cookie := header.Cookie{}

	cookie.Name = "ScrachTest"
	cookie.Value = "deadbeef35"

	//time
	now_time := time.Now()
	add_time := now_time.AddDate(0, 0, 1)
	//Format Function
	dates := fmt.Sprintf("%s", add_time.Format("Mon, 02 Jan 2006 15:04:05 MST"))

	cookie.Expires = dates
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.HttpOnly = "HttpOnly"
	cookie.Secure = "Secure"
	cookie.SameSiteCookie = "Lax"

	return cookie

}
