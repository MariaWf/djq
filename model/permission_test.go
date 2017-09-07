package model

import (
	"testing"
	"fmt"
)

func TestGetPermissionList(t *testing.T) {
	list :=GetPermissionList();
	fmt.Println(list)
	for _,pm := range list{
		fmt.Println(pm.Code)
	}
}
