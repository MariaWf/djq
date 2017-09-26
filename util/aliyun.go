package util

import (
	"strings"
	"path/filepath"
	"mimi/djq/constant"
)

func FormatImageWaterMark(url string) string {
	return formatImage(url, "watermark")
}

func FormatImageShopLogo(url string) string {
	return formatImage(url, "shopLogo")
}

func FormatImagePresent(url string) string {
	return formatImage(url, "present")
}

func FormatImageCashCoupon(url string) string {
	return formatImage(url, "cashCoupon")
}

func FormatImageShopPreImage(url string) string {
	return formatImage(url, "shopPreImage")
}

func FormatImageShopIntroductionImage(url string) string {
	return formatImage(url, "shopIntroductionImage")
}

func FormatImageAdvertisement(url string) string {
	return formatImage(url, "advertisement")
}

func formatImage(url, style string) (newUrl string) {
	newUrl = CleanImageUrl(url)
	fileSuffix := filepath.Ext(newUrl)
	if fileSuffix == "" {
		return
	}
	fileSuffix = strings.ToLower(fileSuffix)
	support := false
	for _, v := range constant.UploadImageSupport {
		if v == fileSuffix && v != ".gif" {
			support = true
		}
	}
	if support {
		newUrl += "?x-oss-process=style/" + style
	}
	return
}

func CleanImageUrl(url string) string {
	i := strings.LastIndex(url, "?")
	if i != -1 {
		return url[:i]
	}
	return url
}