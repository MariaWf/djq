package util

import "testing"

func TestRsaEncrypt(t *testing.T) {
	str := "alsdjf是打發叫是勞動法mimi"
	result, err := RsaEncrypt([]byte(str))
	if err != nil {
		t.Error(err)
	} else {
		source, err := RsaDecrypt(result)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(str)
			t.Log(string(source))
		}
	}
}
