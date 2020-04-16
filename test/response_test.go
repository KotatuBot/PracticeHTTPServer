package test

import (
	"../response"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRead(t *testing.T) {

	var result string
	var expect string
	var filename string

	filename = "test.html"
	result = response.FileRead(filename)
	expect = "<h1>Hello Kotatu</h1>\n<p>Net Work</p>\n"
	assert.Equal(t, result, expect, "Error")

}
