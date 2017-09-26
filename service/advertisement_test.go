package service

import (
	"fmt"
	"mimi/djq/dao/arg"
	"testing"
)

func TestAdvertisement_Get(t *testing.T) {
	serviceObj := &Advertisement{}
	argObj := &arg.Advertisement{}
	argObj.DisplayNames = []string{"name", "link"}
	list, _ := Find(serviceObj, argObj)
	if len(list) > 0 {
		fmt.Println(list[0])
	}
}
