package aliyun

import "testing"

func TestCaptchaSend(t *testing.T) {
	err := CaptchaSend("15017528974", "4141")
	if err != nil {
		t.Error(err)
	}
}
