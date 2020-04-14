package test

import (
	"../parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecvData(t *testing.T) {
	var result string
	var expect string

	bytes := []byte{0x41, 0x41, 0x41, 0x41, 0x41, 0x42, 0x42, 0x00, 0x00}

	result = parse.RecvData(bytes, 3)
	expect = "AAA"
	assert.Equal(t, result, expect, "Error")

}
