package util

import (
	"path/filepath"
	"strings"
)

//func PathAppend(head, tail string) string {
//	head = filepath.ToSlash(head)
//	tail = filepath.ToSlash(tail)
//	hasSlash := false
//	if head != "" && strings.LastIndex(head, "/") == len(head) - 1 {
//		hasSlash = true
//	}
//	if tail != "" && strings.Index(tail, "/") == 0 {
//		if hasSlash {
//			tail = tail[1:]
//		} else {
//			hasSlash = true
//		}
//	}
//	if hasSlash {
//		return head + tail
//	} else {
//		return head + "/" + tail
//	}
//}

func PathAppend(strs ...string) string {
	newStrs := make([]string, 0, len(strs))
	for _, str := range strs {
		str = filepath.ToSlash(str)
		str = strings.Trim(str, "/")
		if str != "" {
			newStrs = append(newStrs, str)
		}
	}
	return strings.Join(newStrs, "/")
}
