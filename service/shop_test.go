package service

import (
	"testing"
	"mimi/djq/dao/arg"
	"fmt"
)

func TestShop_FindByShopClassificationId(t *testing.T) {
	argObj := &arg.Shop{}
	argObj.TargetPage = 2
	argObj.PageSize = 5
	argObj.OrderBy = "priority desc"
	argObj.NotIncludeHide = true
	argObj.DisplayNames = []string{"id", "name", "preImage", "totalCashCouponNumber", "totalCashCouponPrice", "priority"}
	shopClassificationId := "0d91e3dd4f8945c7a82713c3e162d609"
	//shopClassificationId := "11b7e558e0974517a8add2cdffe19a08"
	serviceObj := &Shop{}
	result := serviceObj.FindByShopClassificationId(argObj,shopClassificationId)
	fmt.Println(result.Result)
}
