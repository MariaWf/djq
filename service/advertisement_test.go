package service

import (
	"testing"
	"fmt"
	"mimi/djq/dao/arg"
)

func TestAdvertisement_Get(t *testing.T) {
	serviceObj := &Advertisement{}
	argObj := &arg.Advertisement{}
	list, _ := Find(serviceObj, argObj)
	if len(list) > 0 {
		fmt.Println(list[0])
	}
}
