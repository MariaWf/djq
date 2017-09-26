package util

import "strings"

func IsStringArrEmpty(src []string) bool {
	return src == nil || len(src) == 0
}

func StringArrCopy(list []string) []string {
	newList := make([]string, len(list), len(list))
	copy(newList, list)
	return newList
}

func StringArrClean(list []string) []string {
	newList := make([]string, 0, 10)
	for _, s := range list {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		exist := false
		for _, ns := range newList {
			if ns == s {
				exist = true
				break
			}
		}
		if !exist {
			newList = append(newList, s)
		}
	}
	return newList
}

func StringArrDelete(list []string, obj string) []string {
	if list == nil || len(list) == 0 {
		return make([]string, 0, 10)
	}
	for i, str := range list {
		if str == obj {
			return append(list[:i], list[i+1:]...)
		}
	}
	return StringArrCopy(list)
}

func StringArrConvert2InterfaceArr(list []string) []interface{} {
	if IsStringArrEmpty(list) {
		return nil
	}
	dst := make([]interface{}, len(list), len(list))
	for i, _ := range list {
		dst[i] = list[i]
	}
	return dst
}
