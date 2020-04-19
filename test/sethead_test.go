package test

import (
	"testing"

	"../response"
	"github.com/stretchr/testify/assert"
)

func TestSetStatusCode(t *testing.T) {

	var result string
	var expect string

	expect = "HTTP1.1 200 200"

	result = response.SetStatusCode("200")
	assert.Equal(t, result, expect, "Status Code Error")

}

func TestSetContentEncoding(t *testing.T) {

	var result string
	var expect string
	var data string
	data = "gzip, compress, br"

	result = response.SetContentEncoding(data)
	expect = "gzip"
	assert.Equal(t, result, expect, "Encoding Error")

}

func TestSetContentType(t *testing.T) {
	t.Log("SetContentType")
	t.Run("html", func(t *testing.T) {

		result, _ := response.SetContentType("test.html")
		if result != "text/html;" {
			t.Fatal("text/html Error")

		}

	})

	t.Run("ico", func(t *testing.T) {

		result, _ := response.SetContentType("test.ico")
		if result != "image/x-icon;" {
			t.Fatal("ico Error")

		}

	})

	t.Run("default", func(t *testing.T) {

		result, _ := response.SetContentType("test.txt")
		if result != "text/text;" {
			t.Fatal("default Error")

		}

	})

	t.Log("End SetContentType")

}

func TestSetEtag(t *testing.T) {

	var result string
	var expect string

	var body string

	body = "TEST"
	result = response.SetEtag(body)
	expect = "\"033bd94b1168d7e4f0d644c3c95e35bf\""
	assert.Equal(t, result, expect, "MD5 Hash Error")

}

func TestSetKeepAlive(t *testing.T) {

	expect := "timeout=5"
	expect2 := "max=100"
	result, result2 := response.SetKeepAlive()

	if expect != result && expect2 != result2 {
		t.Error("Not Match keepalive")
	}

}

func TestSetCacheControl(t *testing.T) {
	var result string
	var expect string
	expect = "no-store"

	result = response.SetCacheControl()
	assert.Equal(t, result, expect, "CacheControl Error")

}

func TestXFrameOptions(t *testing.T) {

	t.Run("SAMEORIGIN", func(t *testing.T) {
		result := response.SetXFrameOptions("same")
		expect := "SAMEORIGIN"

		if result != expect {

			t.Fatal("FRAMEOPTION SAME ORIGIN")
		}

	})

	t.Run("DENY", func(t *testing.T) {
		result := response.SetXFrameOptions("deny")
		expect := "DENY"

		if result != expect {

			t.Fatal("FRAMEOPTION DENY")
		}

	})

}
