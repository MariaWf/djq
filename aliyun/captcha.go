package aliyun

import (
	"mimi/djq/config"
)

func CaptchaSend(phoneNumber string, code string) error {
	accessKeyId := config.Get("aliyun_access_key_id")
	accessSecret := config.Get("aliyun_access_secret")
	signName := config.Get("aliyun_sms_sign_name")
	templateCode := config.Get("aliyun_sms_template_code")

	aliyunSms, err := NewAliyunSms(signName, templateCode, accessKeyId, accessSecret)
	if err != nil {
		return err
	}
	err = aliyunSms.Send(phoneNumber, `{"code":"`+code+`"}`)
	if err != nil {
		return err
	}
	return nil
}
