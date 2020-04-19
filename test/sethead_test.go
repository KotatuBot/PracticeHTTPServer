package test

import (
	"testing"

	"../response"
	"github.com/stretchr/testify/assert"
)

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
