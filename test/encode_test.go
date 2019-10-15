package main

import (
	"fmt"
	"github.com/jhonnli/golibs"
	"testing"
)

const (
	_key = "?*-NMn5hJMXoTkm7=dFUYvUJu35UUN_&"
	_iv  = "$8^82_4nc=r045FN"
)

func decordTest(str string) {

	var tmpBytes []byte
	tmpBytes, err := golibs.AesDecrypt(golibs.HexStringToBytes(str), []byte(_key), []byte(_iv))
	if err != nil {
		return
	}
	str = string(tmpBytes)
	println(str)
}

func encordTest(str string) {
	var tmpBytes []byte
	tmpBytes, err := golibs.AesEncrypt([]byte(str), []byte(_key), []byte(_iv))
	if err != nil {
		return
	}
	str = fmt.Sprintf("%x", tmpBytes)
	println(str)
}

func Test_encord(t *testing.T) {
	encordTest("")
}
