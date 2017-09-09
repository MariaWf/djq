package model

import (
	"fmt"
	"testing"
)

func TestGetPermissionList(t *testing.T) {
	list := GetPermissionList()
	fmt.Println(list)
	for _, pm := range list {
		fmt.Println(pm.Code)
	}
}
