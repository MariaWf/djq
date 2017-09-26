package util

import (
	"mimi/djq/model"
	"testing"
)

func TestCleanImageUrl(t *testing.T) {
	url := "http://static.51zxiu.cn/app/djq/upload/image/evidence/8d224a1b041b477fb0b3a80da6186465.jpg?x-oss-process=style/watermark"
	t.Log(FormatImageAdvertisement(url))
	t.Log(CleanImageUrl(url))
}

func TestObj(t *testing.T) {
	a := tk()
	if a == nil {
		t.Log("nil")
	} else {
		t.Log(a)
	}
}

func tk() (admin *model.Admin) {
	if 2 > 1 {
		return
	}
	admin = &model.Admin{}
	return
}
